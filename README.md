# shturmanbot
Пилотный проект - переписываю бота для telegram с PHP на Go. 
Планируемый функционал: 
1. Запоминить введенные пользователем точки для маршрутов (от пунта "А" к пункту "Б").
2. Запоминать доступные средства оплаты проезда (номера карт "Тройка" и/или "Стрелка").
3. Строит оптимальный маршрут с учетом текущего времени и расписания общественного транспорта. 
4. Предлагать маршруты с расписанием на выбор. 
5. Показывать баланс карты "Тройка" и/или "Стрелка" в зависимости от выбранного маршрута.
* Если найдется способ, то пополнять при достижении минимального порога. 
6. Показывать предупреждения о погодных условиях при наличии критических показателей: дождь, сильный ветер, снегопад, резкие перепады температуры.
