import React, { useState } from 'react';
import axios from 'axios';

const OrderSearch = () => {
    const [orderId, setOrderId] = useState('');
    const [order, setOrder] = useState(null);
    const [error, setError] = useState(null);

    const handleSearch = async () => {
        try {
            const response = await axios.get(`http://localhost:4444/api/orders/${orderId}`);
            setOrder(response.data);
            setError(null); // Сброс ошибки при успешном запросе
        } catch (err) {
            setError('Заказ не найден'); // Установите сообщение об ошибке
            setOrder(null); // Сброс результата при ошибке
        }
    };

    return (
        <div>
            <h2>Поиск заказа по ID</h2>
            <input
                type="text"
                value={orderId}
                onChange={(e) => setOrderId(e.target.value)}
                placeholder="Введите ID заказа"
            />
            <button onClick={handleSearch}>Поиск</button>

            {error && <div style={{ color: 'red' }}>{error}</div>}
            {order && (
                <div>
                    <h3>Детали заказа:</h3>
                    <pre>{JSON.stringify(order, null, 2)}</pre>
                </div>
            )}
        </div>
    );
};

export default OrderSearch;