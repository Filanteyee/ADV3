// Слушаем клик на кнопке "Показать все заказы"
document.getElementById("loadOrdersBtn").addEventListener("click", loadOrders);

function loadOrders() {
    // Получаем контейнер для заказов
    const ordersDiv = document.getElementById("ordersContainer");

    // Очищаем контейнер перед загрузкой новых данных
    ordersDiv.innerHTML = "";

    // Загружаем заказы с сервера
    fetch('/orders')
        .then(res => res.json())
        .then(orders => {
            if (orders.length === 0) {
                ordersDiv.innerHTML = "<p>Нет заказов для отображения.</p>";
                return;
            }

            // Добавляем каждый заказ в контейнер
            orders.forEach(order => {
                const orderElement = document.createElement("div");
                orderElement.classList.add("order-card");

                orderElement.innerHTML = `
          <h2>Заказ #${order.id}</h2>
          <p>Статус: ${order.status}</p>
          <div class="order-items">
            ${order.items.map(item => `
              <div class="order-item">
                <img src="${item.image_url}" alt="${item.product_name}" width="100" height="100">
                <p>${item.product_name}</p>
                <p>Цена: ${item.price}</p>
              </div>
            `).join('')}
          </div>
          <div class="order-actions">
            <button class="cancel-btn" onclick="cancelOrder('${order.id}')">Отменить заказ</button>
          </div>
        `;

                ordersDiv.appendChild(orderElement);
            });
        })
        .catch(error => {
            ordersDiv.innerHTML = "<p>Ошибка загрузки заказов. Попробуйте позже.</p>";
        });
}

function cancelOrder(id) {
    if (!confirm("Вы уверены, что хотите отменить заказ?")) return;

    fetch(`/orders/${id}/cancel`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
    })
        .then(res => {
            if (res.ok) {
                alert("Заказ отменён!");
                loadOrders(); // Перезагружаем список заказов
            } else {
                alert("Ошибка при отмене заказа.");
            }
        });
}
