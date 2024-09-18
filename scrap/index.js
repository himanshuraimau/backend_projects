const fs = require('fs');
const { parse } = require('csv-parse');
const { promisify } = require('util');
const csv = require('csv-writer').createObjectCsvWriter;
const puppeteer = require('puppeteer');


const parseAsync = promisify(parse);
const cities = require('./cities.json');

const localityDataPath = 'Locality_village_pincode_final_mar-2017.csv';

let page_limit = 5;
let allData = [];

const getCityLink = (city, page) => {
    return `https://www.makaan.com/${city.toLowerCase().replace(' ', '-')}-residential-property/rent-property-in-${city.toLowerCase().replace(' ', '-')}-city?page=${page}`;
};

async function loadLocalityData() {
    const fileContent = fs.readFileSync(localityDataPath);
    const records = await parseAsync(fileContent, { columns: true });
    return records;
}

async function scrapeCityRent(city, page_limit, localityData) {
    const browser = await puppeteer.launch({ headless: true });
    const page = await browser.newPage();
    let hasNextPage = true;

    try {
        for (let page_number = 1; page_number <= page_limit && hasNextPage; page_number++) {
            const link = getCityLink(city, page_number);
            console.log(`Scraping page ${page_number} for city: ${city}`);
            await page.goto(link, { waitUntil: 'networkidle0' });


            const propertyCards = await page.$$('div[data-type="listing-card"]');

            for (const card of propertyCards) {
                try {
                    const rentInfo = await card.evaluate(() => {
                        const localityElement = document.querySelector('.locName .loclink');
                        const locality = localityElement?.querySelector('span[itemprop="addressLocality"] strong')?.textContent.trim() || 'N/A';
                        const city = localityElement?.querySelector('.cityName')?.textContent.trim() || 'N/A';

                        const rentElement = document.querySelector('.price .val');
                        const rent = rentElement?.textContent.trim() || 'N/A';

                        return {
                            locality: locality,
                            city: city,
                            rent: rent
                        };
                    });


                    const localityMatch = localityData.find(record => record['Village/Locality name'].toLowerCase() === rentInfo.locality.toLowerCase());
                    const pincode = localityMatch ? localityMatch.Pincode : 'N/A';

                    allData.push({
                        locality: rentInfo.locality,
                        city: rentInfo.city,
                        rent: rentInfo.rent,
                        pincode: pincode
                    });

                    console.log(`Scraped: ${rentInfo.locality}, ${rentInfo.city} - Rent: ${rentInfo.rent}, Pincode: ${pincode}`);

                } catch (error) {
                    console.error('Error extracting data from property card:', error);
                }
            }


            const nextPageButton = await page.$('[aria-label="nextPage"]');
            if (!nextPageButton || page_number === page_limit) {
                hasNextPage = false;
            }
        }
    } catch (error) {
        console.error('Error during scraping:', error);
    } finally {
        await browser.close();
    }
}


function saveToCSV() {
    const csvWriter = csv({
        path: `all_cities_rent_data.csv`,
        header: [
            { id: 'locality', title: 'Locality' },
            { id: 'city', title: 'City' },
            { id: 'rent', title: 'Rent' },
            { id: 'pincode', title: 'Pincode' }
        ]
    });

    csvWriter.writeRecords(allData)
        .then(() => console.log(`All data has been written to all_cities_rent_data.csv successfully`));
}

async function scrapeAllCities() {
    const localityData = await loadLocalityData();

    for (const cityObj of cities) {
        const cityName = cityObj.name;
        console.log(`Scraping data for city: ${cityName}`);
        await scrapeCityRent(cityName, page_limit, localityData);
    }
    saveToCSV();  
}

scrapeAllCities();
