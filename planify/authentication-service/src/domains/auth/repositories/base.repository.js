class BaseRepository {
    constructor(model) {
        this.model = model;
    }

    async findById(id) {
        return this.model.findByPk(id);
    }

    async findOne(conditions) {
        return this.model.findOne(conditions);
    }

    async findAll(conditions = {}) {
        return this.model.findAll(conditions);
    }

    async create(data) {
        return this.model.create(data);
    }

    async update(id, data) {
        const entity = await this.findById(id);
        if (!entity) return null;
        return entity.update(data);
    }

    async delete(id) {
        const entity = await this.findById(id);
        if (!entity) return false;
        await entity.destroy();
        return true;
    }
}

module.exports = BaseRepository;
