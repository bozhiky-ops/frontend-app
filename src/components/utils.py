# utils.py
from typing import List, Dict, Tuple
from datetime import datetime
import os
import hashlib
import logging

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

def hash_string(input_str: str) -> str:
    """Hash a string using SHA-256."""
    return hashlib.sha256(input_str.encode()).hexdigest()

def get_current_timestamp() -> float:
    """Get the current timestamp in seconds since the epoch."""
    return datetime.now().timestamp()

def is_valid_url(url: str) -> bool:
    """Check if a URL is valid."""
    try:
        result = urllib.parse.urlparse(url)
        return all([result.scheme, result.netloc])
    except ValueError:
        return False

def get_file_size(file_path: str) -> int:
    """Get the size of a file in bytes."""
    return os.path.getsize(file_path)

def get_file_extension(file_path: str) -> str:
    """Get the file extension from a file path."""
    return os.path.splitext(file_path)[1].lstrip('.')

def load_json_file(file_path: str) -> Dict:
    """Load a JSON file and return its contents as a dictionary."""
    with open(file_path, 'r') as file:
        return json.load(file)

def save_json_file(file_path: str, data: Dict) -> None:
    """Save a dictionary to a JSON file."""
    with open(file_path, 'w') as file:
        json.dump(data, file, indent=4)

def get_directory_contents(directory_path: str) -> List:
    """Get a list of files and subdirectories in a directory."""
    return os.listdir(directory_path)

def get_file_path(directory_path: str, file_name: str) -> str:
    """Get the full path of a file in a directory."""
    return os.path.join(directory_path, file_name)

def is_binary_file(file_path: str) -> bool:
    """Check if a file is a binary file."""
    with open(file_path, 'rb') as file:
        return not file.read(1024).decode('utf-8').strip()

import urllib.parse
import json