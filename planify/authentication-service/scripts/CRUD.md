To use the generator:

1. Install dependencies:
```bash
npm install commander lodash
```

2. Make the script executable:
```bash
chmod +x scripts/generate-crud.js
```

3. Run the generator:
```bash
# Basic usage
./scripts/generate-crud.js user

# With fields
./scripts/generate-crud.js product --fields "name:string,price:decimal,description:text,category:string"
```

The generator will create:
- Complete directory structure
- Model with specified fields
- Controller with CRUD operations
- Service layer with business logic
- Repository for data access
- DTOs for request/response
- Validation rules
- Routes configuration
- Swagger documentation
- Database migration

Example usage:
```bash
./scripts/generate-crud.js product --fields "name:string,price:decimal,description:text,category:string"
```

This will generate:
```
src/domains/product/
├── controllers/
│   └── product.controller.js
├── services/
│   └── product.service.js
├── repositories/
│   └── product.repository.js
├── models/
│   └── product.model.js
├── dtos/
│   └── product.dto.js
├── validations/
│   └── product.validation.js
├── routes/
│   └── product.routes.js
└── docs/
    └── product.swagger.js

migrations/
└── YYYYMMDDHHMMSS-create-product.js
```
