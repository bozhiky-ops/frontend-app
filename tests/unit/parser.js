const fs = require('fs');
const path = require('path');

class Parser {
    constructor(filePath) {
        this.filePath = path.resolve(filePath);
    }

    parse() {
        return new Promise((resolve, reject) => {
            fs.readFile(this.filePath, 'utf8', (err, data) => {
                if (err) {
                    reject(`Error reading file: ${err.message}`);
                } else {
                    try {
                        const parsedData = JSON.parse(data);
                        resolve(parsedData);
                    } catch (parseError) {
                        reject(`Error parsing JSON: ${parseError.message}`);
                    }
                }
            });
        });
    }

    static validateSchema(data, schema) {
        const requiredFields = schema.required || [];
        for (const field of requiredFields) {
            if (!data.hasOwnProperty(field)) {
                throw new Error(`Missing required field: ${field}`);
            }
        }
        return true;
    }
}

module.exports = Parser;