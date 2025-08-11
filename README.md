# Resume builder

This repo is a templated way for me to build and maintain my resume.

## Resume Guidelines I'm following

Adhering to the resume format from this reddit subthread: [r/EngineeringResumes](https://www.reddit.com/r/EngineeringResumes/)

Here is their Wiki: [r/EngineeringResumes/wiki/index/](https://www.reddit.com/r/EngineeringResumes/wiki/index/)

This is based on their "template": [EngineeringResumes Template (Google Doc Link)](https://old.reddit.com/r/EngineeringResumes/wiki/templates)

## Concept
- Define resume data within a .yml file
- Extract it and consume it in go templates
- Compile it to an output.html file
- Print it to PDF

## Technologies
- Go
  - used to read yml
  - send data to go templates
  - take the output of those templates combined into a single .html file

- html
  - structuring the resume data into a neat presentable format

- vanilla css

