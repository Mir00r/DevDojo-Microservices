#!/usr/bin/env node
const fs = require('fs');
const path = require('path');
const {program} = require('commander');
const {camelCase, upperFirst, snakeCase, upperCase} = require('lodash');

// Config for base paths
const BASE_PATH = path.join(process.cwd(), 'src/domains/auth');
const TEMPLATE_PATH = path.join(__dirname, 'templates');

program
    .name('generate-crud')
    .description('Generate CRUD files for a new domain')
    .argument('<domain>', 'Domain name (e.g., user, product)')
    .option('-f, --fields <fields>', 'Fields in format: name:type,email:string,age:integer')
    .parse(process.argv);

const domainName = program.args[0];
const fields = program.opts().fields?.split(',') || [];

// Parse fields into structured format
const parseFields = (fields) => {
    return fields.map(field => {
        const [name, type] = field.split(':');
        return {name, type: type || 'string'};
    });
};

const parsedFields = parseFields(fields);

// Template replacement function
const replaceTemplateVars = (template, vars) => {
    return template.replace(/\{\{(\w+)\}\}/g, (match, key) => vars[key] || match);
};

// Generate directory structure
const createDirectoryStructure = (domainName) => {
    const domainPath = path.join(BASE_PATH, domainName);
    const directories = [
        '',
        'controllers',
        'services',
        'repositories',
        'models',
        'dtos',
        'validations',
        'routes',
        'docs'
    ];

    directories.forEach(dir => {
        const fullPath = path.join(domainPath, dir);
        if (!fs.existsSync(fullPath)) {
            fs.mkdirSync(fullPath, {recursive: true});
        }
    });

    return domainPath;
};

// Generate files based on templates
const generateFiles = (domainPath, domainName, fields) => {
    const templateVars = {
        DomainName: upperFirst(camelCase(domainName)),
        domainName: camelCase(domainName),
        DOMAIN_NAME: upperCase(snakeCase(domainName)),
        fields: fields
    };

    // Generate model
    generateFile('model.template.js',
        path.join(domainPath, 'models', `${domainName}.model.js`),
        templateVars
    );

    // Generate controller
    generateFile('controller.template.js',
        path.join(domainPath, 'controllers', `${domainName}.controller.js`),
        templateVars
    );

    // Generate service
    generateFile('service.template.js',
        path.join(domainPath, 'services', `${domainName}.service.js`),
        templateVars
    );

    // Generate repository
    generateFile('repository.template.js',
        path.join(domainPath, 'repositories', `${domainName}.repository.js`),
        templateVars
    );

    // Generate DTOs
    generateFile('dto.template.js',
        path.join(domainPath, 'dtos', `${domainName}.dto.js`),
        templateVars
    );

    // Generate validation
    generateFile('validation.template.js',
        path.join(domainPath, 'validations', `${domainName}.validation.js`),
        templateVars
    );

    // Generate routes
    generateFile('routes.template.js',
        path.join(domainPath, 'routes', `${domainName}.routes.js`),
        templateVars
    );

    // Generate Swagger docs
    generateFile('swagger.template.js',
        path.join(domainPath, 'docs', `${domainName}.swagger.js`),
        templateVars
    );

    // Generate migration
    const timestamp = new Date().toISOString().replace(/\D/g, '').slice(0, 14);
    generateFile('migration.template.js',
        path.join(process.cwd(), 'migrations', `${timestamp}-create-${snakeCase(domainName)}.js`),
        templateVars
    );
};

// Helper function to generate a single file
const generateFile = (templateName, outputPath, vars) => {
    const templateContent = fs.readFileSync(
        path.join(TEMPLATE_PATH, templateName),
        'utf8'
    );
    const content = replaceTemplateVars(templateContent, vars);
    fs.writeFileSync(outputPath, content);
    console.log(`Generated: ${outputPath}`);
};

// Main execution
try {
    const domainPath = createDirectoryStructure(domainName);
    generateFiles(domainPath, domainName, parsedFields);
    console.log(`\n✅ Successfully generated CRUD files for ${domainName}`);
} catch (error) {
    console.error('❌ Error generating files:', error);
    process.exit(1);
}
