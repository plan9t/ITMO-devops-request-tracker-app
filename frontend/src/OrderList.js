import React, { useEffect, useState } from 'react';
import axios from 'axios';

const OrderList = () => {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchOrders = async () => {
            try {
                const response = await axios.get('http://localhost:4444/api/orders');
                setOrders(response.data);
            } catch (err) {
                setError(err);
            } finally {
                setLoading(false);
            }
        };

        fetchOrders();
    }, []);

    if (loading) return <div>Загрузка...</div>;
    if (error) return <div>Ошибка: {error.message}</div>;

    return (
        <div>
            <h1>Список заказов</h1>
            <ul>
                {orders.map(order => (
                    <li key={order.order_uid}>{order.track_number} - {order.entry}</li>
                ))}
            </ul>
        </div>
    );
};

export default OrderList;