# **Database Migration Guide for PostgreSQL with Sequelize**

This guide provides detailed instructions on how to manage database migrations using **Sequelize CLI** for a PostgreSQL database with **multiple schemas**.

---

## **📌 Prerequisites**
Before running migrations, ensure that:
- You have **Node.js** and **npm/yarn** installed.
- You have **PostgreSQL** installed and running.
- You have configured Sequelize and database settings in `config/config.js`.
- You have installed Sequelize CLI:
  ```sh
  npm install --save-dev sequelize-cli
  ```


## **📌 Initialize Sequelize**
Run the following command to initialize Sequelize in your project:

 ```sh
  npx sequelize-cli init
  ```
This will create the following structure:

```
│── migrations/        # Database migration files
│── models/            # ORM Models
│── seeders/           # Seed data (optional)
│── config/            # Sequelize configuration
```

---

## **🔹 Basic Migration Commands**

### **1️⃣ Create a New Migration File**
To create a new migration file, use:
```sh
npx sequelize-cli migration:generate --name create-users-table
```
This creates a file under `migrations/`:
```
migrations/
│── 20240205010101-create-users-table.js
```
Now, you can define the table schema inside this file.

---

### **2️⃣ Run Migrations**
To apply all pending migrations, run:
```sh
npx sequelize-cli db:migrate
```
This will create the tables based on migration files in the database.

✅ **Example Output:**
```
== 20240205010101-create-users-table: migrated (SUCCESS)
```

---

### **3️⃣ Rollback (Undo) Migrations**
If you need to undo the last migration, run:
```sh
npx sequelize-cli db:migrate:undo
```
If you need to rollback **all** migrations:
```sh
npx sequelize-cli db:migrate:undo:all
```

✅ **Example Output:**
```
== 20240205010101-create-users-table: reverted
```

---

## **🔹 Managing Schemas in PostgreSQL**
In PostgreSQL, we use **schemas** (like `auth`, `inventory`, etc.). Sequelize supports schemas using `schema: 'your_schema_name'` in migration files.

### **Example Migration with Schema (`auth` Schema)**
#### **1️⃣ Create `roles` Table First**
```js
module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.createTable(
      'roles',
      {
        id: { type: Sequelize.INTEGER, autoIncrement: true, primaryKey: true },
        name: { type: Sequelize.STRING, allowNull: false },
        createdAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
        updatedAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
      },
      { schema: 'auth' } // Define Schema Here
    );
  },
  down: async (queryInterface) => {
    await queryInterface.dropTable('roles', { schema: 'auth' });
  },
};
```

#### **2️⃣ Create `users` Table After `roles`**
```js
module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.createTable(
      'users',
      {
        id: { type: Sequelize.INTEGER, autoIncrement: true, primaryKey: true },
        email: { type: Sequelize.STRING, allowNull: false, unique: true },
        password: { type: Sequelize.STRING, allowNull: false },
        roleId: {
          type: Sequelize.INTEGER,
          references: { model: 'roles', key: 'id', schema: 'auth' }, // Reference with Schema
          onDelete: 'CASCADE',
        },
        createdAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
        updatedAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
      },
      { schema: 'auth' } // Define Schema Here
    );
  },
  down: async (queryInterface) => {
    await queryInterface.dropTable('users', { schema: 'auth' });
  },
};
```

---

## **🔹 How to Add a New Table (Entity)**
To add a new entity, follow these steps:

1️⃣ **Generate a Migration File**
```sh
npx sequelize-cli migration:generate --name create-products-table
```

2️⃣ **Edit the Migration File**  
Modify `up()` and `down()` methods with schema support:
```js
module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.createTable(
      'products',
      {
        id: { type: Sequelize.INTEGER, autoIncrement: true, primaryKey: true },
        name: { type: Sequelize.STRING, allowNull: false },
        price: { type: Sequelize.FLOAT, allowNull: false },
        createdAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
        updatedAt: { type: Sequelize.DATE, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
      },
      { schema: 'inventory' } // Define Schema Here
    );
  },
  down: async (queryInterface) => {
    await queryInterface.dropTable('products', { schema: 'inventory' });
  },
};
```

3️⃣ **Run the Migration**
```sh
npx sequelize-cli db:migrate
```

---

## **🔹 How to Modify an Existing Table**
If you need to **add a new column** (e.g., `phoneNumber` in `users` table):

1️⃣ **Generate Migration File**
```sh
npx sequelize-cli migration:generate --name add-phoneNumber-to-users
```

2️⃣ **Edit the Migration File**
```js
module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.addColumn('users', 'phoneNumber', { type: Sequelize.STRING }, { schema: 'auth' });
  },
  down: async (queryInterface) => {
    await queryInterface.removeColumn('users', 'phoneNumber', { schema: 'auth' });
  },
};
```

3️⃣ **Run Migration**
```sh
npx sequelize-cli db:migrate
```

---

## **🔹 How to Rename or Remove a Column**
1️⃣ **Generate Migration File**
```sh
npx sequelize-cli migration:generate --name rename-user-column
```

2️⃣ **Modify the Migration**
```js
module.exports = {
  up: async (queryInterface) => {
    await queryInterface.renameColumn({ tableName: 'users', schema: 'auth' }, 'phoneNumber', 'contactNumber');
  },
  down: async (queryInterface) => {
    await queryInterface.renameColumn({ tableName: 'users', schema: 'auth' }, 'contactNumber', 'phoneNumber');
  },
};
```

3️⃣ **Run Migration**
```sh
npx sequelize-cli db:migrate
```

---

## **🔹 How to Seed Initial Data**
To insert default roles:
```sh
npx sequelize-cli seed:generate --name seed-roles
```

Edit the `seeders/XXXXXX-seed-roles.js` file:
```js
module.exports = {
  up: async (queryInterface) => {
    await queryInterface.bulkInsert(
      { tableName: 'roles', schema: 'auth' },
      [{ id: 1, name: 'Admin' }, { id: 2, name: 'User' }],
      {}
    );
  },
  down: async (queryInterface) => {
    await queryInterface.bulkDelete({ tableName: 'roles', schema: 'auth' }, null, {});
  },
};
```

Run the seed:
```sh
npx sequelize-cli db:seed:all
```

---

## **📌 Summary Table**
| Command                                                    | Purpose |
|------------------------------------------------------------|---------|
| `npx sequelize-cli migration:generate --name create-table` | Generate a new migration file |
| `npx sequelize-cli db:migrate`                             | Run all migrations |
| `npx sequelize-cli db:migrate:undo`                        | Undo the last migration |
| `npx sequelize-cli db:migrate:undo:all`                    | Undo all migrations |
| `npx sequelize-cli seed:generate --name seed-data`         | Generate a seeder file |
| `npx sequelize-cli db:seed:all`                            | Run all seeders |
| `npx sequelize-cli init`                                   | initialize Sequelize in your project |

---

## **✅ Conclusion**
- Always ensure migration **order is correct** (e.g., create `roles` before `users`).
- Use **schemas** to organize tables logically (`auth`, `inventory`, etc.).
- **Rollback carefully** if a migration fails.
- Always test changes **in a development environment** before running on production.

🔹 Following these guidelines ensures **smooth database version control**! 🚀
