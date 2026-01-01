// parser.js

/**
 * Module dependencies.
 */
const fs = require('fs');
const path = require('path');
const { console } = require('console');

/**
 * File parser module.
 */
class FileParser {
  /**
   * Reads a file and returns its contents.
   *
   * @param {string} filePath - Path to the file to read.
   * @returns {string} File contents.
   */
  readFileSync(filePath) {
    try {
      return fs.readFileSync(filePath, 'utf8');
    } catch (error) {
      console.error(`Error reading file: ${error.message}`);
      process.exit(1);
    }
  }

  /**
   * Parses a file into an array of lines.
   *
   * @param {string} filePath - Path to the file to parse.
   * @returns {string[]} File contents as an array of lines.
   */
  parseFileSync(filePath) {
    const fileContents = this.readFileSync(filePath);
    return fileContents.split('\n').filter(line => line.trim());
  }
}

module.exports = FileParser;