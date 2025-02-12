class BaseDto {
    static from(data) {
        return new this(data);
    }

    static validate(data) {
        const errors = [];
        if (!data) {
            errors.push('Data is required');
        }
        return errors;
    }
}

module.exports = BaseDto;
