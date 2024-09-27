module.exports = {
    transform: {
        '^.+\\.jsx?$': 'babel-jest', // Для обработки .js и .jsx файлов
    },
    testEnvironment: 'jsdom', // Используйте jsdom для тестирования React компонентов
    moduleFileExtensions: ['js', 'jsx', 'json', 'node'],
};