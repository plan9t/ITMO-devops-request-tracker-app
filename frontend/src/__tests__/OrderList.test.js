// OrderList.test.js
import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom'; // Импортируем jest-dom для toBeInTheDocument
import OrderList from '../OrderList';

// Мокаем глобальный fetch
beforeEach(() => {
    global.fetch = jest.fn();
});

afterEach(() => {
    jest.clearAllMocks(); // Очищаем моки после каждого теста
});

describe('OrderList component', () => {
    test('renders loading state initially', () => {
        render(<OrderList />);

        // Проверяем, что отображается текст "Загрузка..."
        expect(screen.getByText(/загрузка/i)).toBeInTheDocument();
    });

    test('renders orders when fetch is successful', async () => {
        const mockOrders = [
            { order_uid: '1', track_number: 'TRACK123', entry: 'Entry 1' },
            { order_uid: '2', track_number: 'TRACK456', entry: 'Entry 2' },
        ];

        // Настройка мока для fetch
        global.fetch.mockResolvedValueOnce({
            ok: true,
            json: jest.fn().mockResolvedValueOnce(mockOrders),
        });

        render(<OrderList />);

        // Ожидание, пока список заказов будет отрендерен
        await waitFor(() => expect(screen.getByText(/список заказов/i)).toBeInTheDocument());

        // Проверка, что заказы отображаются на экране
        expect(screen.getByText('TRACK123 - Entry 1')).toBeInTheDocument();
        expect(screen.getByText('TRACK456 - Entry 2')).toBeInTheDocument();
    });

    test('renders error message when fetch fails', async () => {
        // Настройка мока для fetch, чтобы он вызывал ошибку
        global.fetch.mockRejectedValueOnce(new Error('Ошибка загрузки'));

        render(<OrderList />);

        // Ожидание, пока ошибка будет отрендерена
        await waitFor(() => expect(screen.getByText(/ошибка:/i)).toBeInTheDocument());

        // Проверка, что сообщение об ошибке отображается на экране
        expect(screen.getByText(/ошибка загрузки/i)).toBeInTheDocument();
    });
});