const {Op} = require('sequelize');

/**
 * Base repository class for common database operations
 */
class BaseRepository {
    /**
     * @param {Model} model - Sequelize model
     */
    constructor(model) {
        this.model = model;
    }

    /**
     * Find entity by primary key
     * @param {number|string} id - Primary key
     * @param {Object} options - Additional options (include, attributes, etc.)
     * @returns {Promise<Model>}
     */
    async findById(id, options = {}) {
        return this.model.findByPk(id, options);
    }

    /**
     * Find single entity by conditions
     * @param {Object} options - Sequelize query options
     * @returns {Promise<Model>}
     */
    async findOne(options = {}) {
        return this.model.findOne(options);
    }

    /**
     * Find all entities matching conditions
     * @param {Object} options - Sequelize query options
     * @returns {Promise<Model[]>}
     */
    async findAll(options = {}) {
        return this.model.findAll(options);
    }

    /**
     * Find and count all entities with pagination
     * @param {Object} options - Sequelize query options
     * @returns {Promise<Object>} - { count, rows }
     */
    async findAndCountAll(options = {}) {
        return this.model.findAndCountAll(options);
    }

    /**
     * Create new entity
     * @param {Object} data - Entity data
     * @param {Object} options - Creation options (transaction, etc.)
     * @returns {Promise<Model>}
     */
    async create(data, options = {}) {
        return this.model.create(data, options);
    }

    /**
     * Update entity by id
     * @param {number|string} id - Entity id
     * @param {Object} data - Update data
     * @param {Object} options - Update options
     * @returns {Promise<Model>}
     */
    async update(id, data, options = {}) {
        const entity = await this.findById(id);
        if (!entity) return null;
        return entity.update(data, options);
    }

    /**
     * Delete entity by id
     * @param {number|string} id - Entity id
     * @param {Object} options - Delete options
     * @returns {Promise<boolean>}
     */
    async delete(id, options = {}) {
        const entity = await this.findById(id);
        if (!entity) return false;
        await entity.destroy(options);
        return true;
    }

    /**
     * Count entities matching conditions
     * @param {Object} options - Count options
     * @returns {Promise<number>}
     */
    async count(options = {}) {
        return this.model.count(options);
    }

    /**
     * Bulk create entities
     * @param {Object[]} data - Array of entity data
     * @param {Object} options - Bulk create options
     * @returns {Promise<Model[]>}
     */
    async bulkCreate(data, options = {}) {
        return this.model.bulkCreate(data, options);
    }

    /**
     * Bulk update entities
     * @param {Object} data - Update data
     * @param {Object} options - Update options with where clause
     * @returns {Promise<[number, Model[]]>}
     */
    async bulkUpdate(data, options = {}) {
        return this.model.update(data, options);
    }
}

module.exports = BaseRepository;
