document.getElementById('calculator-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    console.log("Форма отправлена");

    const expression = document.getElementById('expression').value;
    const resultDiv = document.getElementById('result');
    resultDiv.textContent = 'Вычисление...';

    try {
        const response = await fetch('/api/v1/calculate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ expression }),
        });

        console.log("Ответ сервера:", response);

        if (!response.ok) {
            throw new Error('Ошибка при отправке выражения');
        }

        const data = await response.json();
        console.log("Полученные данные:", data);

        resultDiv.textContent = `Результат выражения: ${data.result}`;
    } catch (error) {
        console.error(error);
        resultDiv.textContent = `Ошибка: ${error.message}`;
    }
});

document.getElementById('settings-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    const settings = {
        timeAddition: parseInt(document.getElementById('time-addition').value),
        timeSubtraction: parseInt(document.getElementById('time-subtraction').value),
        timeMultiplication: parseInt(document.getElementById('time-multiplication').value),
        timeDivision: parseInt(document.getElementById('time-division').value),
    };

    try {
        const response = await fetch('http://localhost:8080/api/v1/settings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(settings),
        });

        if (!response.ok) {
            throw new Error('Ошибка при сохранении настроек');
        }

        alert('Настройки успешно сохранены!');
    } catch (error) {
        alert(`Ошибка: ${error.message}`);
    }
});