const getPaginationParams = (query) => {
    const page = parseInt(query.page, 10) || 1;
    const limit = parseInt(query.limit, 10) || 10;
    const offset = (page - 1) * limit;

    return {
        limit,
        offset,
        page
    };
};

const formatPaginatedResponse = (data, page, limit) => {
    const {count: totalItems, rows: items} = data;
    const currentPage = page;
    const totalPages = Math.ceil(totalItems / limit);

    return {
        items,
        meta: {
            totalItems,
            itemsPerPage: limit,
            totalPages,
            currentPage
        }
    };
};

module.exports = {
    getPaginationParams,
    formatPaginatedResponse
};
