



function getSize(inputSize) {
    const sizeMap = {
        s: '80px',
        m: '120px',
        l: '160px',
        xl: '200px',
        xxl: '240px'
    };

    return sizeMap[inputSize] || '120px';
}
