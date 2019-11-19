# Адитья Бхаргава. Грокаем алгоритмы. Конспект книги с примерами реализации алгоритмов на Go

## Глава 1. Знакомство с алгоритмами

### Бинарный поиск
Алгоритм на входе получает отсортированный список элементов. Возвращает позицию искомого элемента или null.

Пример: игра "угадай число". Загадано число от 1 до 100. При каждой попытке возвращается ответ: "мало", "много" или "угадал".
Плохой способ – перебирать все числа подряд. При самом отрицательном сценарии потребуется 100 попыток.
Эффективный способ – начнем с 50. Если мало, то пробуем 75. Если много, то 63 и тд, каждый раз исключая половину оставшихся возможных вариантов.
Какое бы число не было задумано, его можно угадать не более чем за 7 попыток.

При простом поиске из 240 000 вариантов может потребоваться 240 000 попыток. При бинарном - максимум 18.
Для списка из n элементов простой поиск занимает n шагов, бинарный log2n шагов.

*Логарифм по смыслу противоположен возведению в степень. Логарифм по основанию 10 от 100 означает в какую степень нужно возвести 10, чтобы получить 100. Ответ 10*

В нотации "О большое"(об этом позже), log всегда означает 2. Для списка из 8 элементов log8 == 3, тк 2^3 === 8.

**Бинарный поиск работает только с отсортированным списком.**


### Время выполнения
Скорость измеряется не временем, а ростом кол-ва операций.

Если максимальное к-во попыток совпадает с размером списка, при простом поиске,  время выполнения - линейное время пыполнения.

При бинарном поиске - поиск выполняется за логарифмическое время.

Факториальное время - ужасный ужас - при очень длинном списке, поиск становится невероятно долгим. Наример, поиск кратчайшего расстояния для того, чтобы объехать n городов:
5 городов - 120 операций, 6 городов - 720, 7 городов - 5040.

**линейное время - O(n)**

**логарифмическое время - O(log n)**

**факториальное время - O(n!)**

[Реализация бинарного поиска](chapter_01/binary_search.go)

[Упражнения](chapter_01/binarysearch_tasks.md)

## Глава 2. Сортировка выбором

### Как устроена память
Память компьютера можно представить в виде огромного шкафа с множеством ящиков. У каждого ящика есть адрес. 
Когда требуется сохранить в памяти какое-то значение, мы запрашиваем у компьютера место в памяти, а он выдает адрес для сохранения значения.
Если нужно сохранить несколько значений, есть 2 основных способа: массивы и связные списки.

### Массивы и связные списки
При использовании массива, все задачи хранятся в памяти непрерывно, друг за другом. 
В списках каждый элемент хранит в себе ссылку на следующий.
При этом элементы могут размещаться в памяти где угодно. 

Получить запись быстрее можно из массива - обратиться по нужному адресу. При такой же операции со списком нужно пройти по всей цепочке.
С записью - наоборот. Нужна всего одна операция для записи значения в список и n операций для записи в массив.

чтение из списка – O(n) запись в список – O(1) удаление из списка – O(1)

чтение из массива – O(1) запись в массив – O(n) удаление из массива – O(n)

### Сортировка выбором
Легкий для понимания алгоритм, но очень медленно работает – O(n^2);
Каждый раз перебирается n элементов-1

[Реализация сортировки выбором](chapter_02/sort_by_selection.go)

[Упражнения](chapter_02/sortbyselection_tasks.md)


## Глава 3. Рекурсия
Рекурсия - вызов функции самой себя. Состоит из двух частей: базового и рекурсивного случая. В базовом - условие выхода.

### Стек вызовов
Под вызов каждой функции определяется блок в памяти для всех переменных. Блоки определяются в стек.
Вызов ф-и из другой ф-и приостанавливает ее действие - частично завершенное состояние.

[Пример рекурсии при вычислении факториала](chapter_03/factorial.go)

[Пример рекурсии обратного отсчёта](chapter_03/countdown.go)

[Упражнения](chapter_03/recursion_tasks.md)

## Глава 4. Быстрая сортировка
Работает быстрее сортировки выбором и часто применяется. Базовый случай - пустой массив или массив из одного элемента.
O(n log n) – в среднем. В худшем случае – O(n^2); Зависит от опорного элемента.

[Реализация быстрой сортировки](chapter_04/quicksort.go)

[Упражнения](chapter_04/recursion.go)

[Бенчмарки](chapter_04/bench_test.go)

## Глава 5. Хеш-таблица
Хэш-функция – функция, которая получает строку(набор байтов) и возвращает число. Хэш-таблицы – структура данных, связывающая ключи со значениями.
O(1) – в среднем. Отлично подходят для хранения кэша. 

### Коллизии
Коллизия - ситуация, когда двум ключам назначается один элемент массива. Простейшее решение – связный список в этом элементе. Хорошая хэш-функция создает минимальное кол-во коллизий.

[Упражнения](chapter_05/hash_tasks.md)

## Глава 6. Поиск в ширину

### Знакомство с графами
Граф моделирует набор связей. Каждый граф состоит из узлов и ребер. 

### Поиск в ширину
Этот алгоритм отвечает на два вопроса: существует ли путь от одного узла к другому и какой путь кратчайший.
Поиск производится сначала по связям первого уровня, потом второго урвня и тд. Связи, по которым нужно продолжать поиск добавляются в список
и образуют очередь. Очередь относится к категории структур FIFO (First In, First Out).

Время выполнения – O(V+E) V – кол-во вершин, E – кол-во ребер;

Дерево - граф, в котором нет ребер, указывающих в обратном направлении

[Реализация поиска в ширину](chapter_06/wide_search.go)

[Упражнения](chapter_06/widesearch_tasks.md)


## Глава 7. Алгоритм Дейкстры
Если известна "стоимость" каждого ребра, граф – взвешенный. Для поиска кратчайшего пути во взвешенном(без отрицательных весов), направленном графе, используется алгоритм Дейкстры.
Хранение данных графа реализовал в двух вариантах. Первый (медленный и затратный) в виде списков структур. Второй вариант в виде хэш-таблиц.
Стоимость узлов расчитывается относительно любого узла заданного стартовым. При указании необязательного конечного узла, во втором варианте программы, будет нарисован кратчайший путь.

[Реализация алгоритма Дейкстры V1](chapter_07/dijckstra_v1.go)

[Реализация алгоритма Дейкстры V2](chapter_07/dijckstra_v2.go)

[Упражнения](chapter_07/dijckstra_tasks.md)

[Бенчмарки](chapter_07/bench_test.go)

Бенчмарками этих функций проверял разницу в работе со списком и с хэш-таблицей.
[Функции для бенчмарков](chapter_07/func_for_bench.go)

## Глава 8. Жадные алгоритмы
На каждом шаге алгоритма выбирается локальное оптимальное решение. В итоге получается глобально-оптимальное.
Время выполнения – O(n^2)
Эти алгоритмы являются приближенными. Они применяются когда время вычисления точного решения занимает слишком много времени. Эффективность такого алгоритма оценивается по:
* быстроте
* близости полученного результата к оптимальному

В задаче о коммивояжере зависимость вариантов от числа городов является факториальной. И эта задача и задача о покрытии множества являются NP-полными. Это задачи для которых невозможно написать быстрый точный алгоритм. Они решаются только приближенно.
Нет однозначного способа определения, что задача является NP-полной.
Вот несколько характерных признаков:
* алгоритм быстро работает при малом количестве элементов, но сильно замедляется при их увеличении
* формулировка _"все комбинации Х"_ часто указывает на NP- полноту задачи
* задачу невозможно разбить на мелкие подзадачи и приходиться вычислять все возможные варианты Х
* если в задаче встречается некоторая последовательность и задача не имеет простого решения
* если в задаче встречается некоторое множество и задача не имеет простого решения
* если задачу можно переформулировать в условия задачи покрытия множества или в задачу о коммивояжере

[Реализация жадного алгоритма](chapter_08/greedy.go)

[Упражнения](chapter_08/greedy_tasks.md)

## Глава 9. Динамическое программирование
Динамическое программирование - это метод решения сложных задач, разбиваемых на подзадачи, которые решаются в первую очередь.
Динамическое программирование работате только в том случае, если каждая подзадача автономна, то есть не зависит от других подзадач.
Применяется для оптимизации какой-либо характеристики при заданных ограничениях.
Общие рекомендации по решениям:
* в каждом решении из области динамического программирования строится таблица
* значения ячеек таблицы обычно соответствует оптимизируемой характеристике
* каждая ячейка представляет подзадачу, поэтому нужно думать, как разбить задачу на подзадачи. Это поможет определиться с осями

[Реализация динамического программирования](chapter_09/dynamic.go)

[Упражнения](chapter_09/dynamic_tasks.md)

## Глава 10. Алгоритм k-ближайших соседей
При классификации неизвестного надо оценить характеристики его ближайших соседей, если они похожи - неизвестное скорее всего является одним из оцененных. 
Характеристики или признаки по которым ведется сравнение можно представить в виде координат между которыми вычисляется расстояние по формуле _Пифагора_:
 L=SQRT((x1-x2)^2 + (y1-y2)^2 +...+(n1-n2)^2)
Координат может быть любое количество, формула не изменится
Чем меньше расстояние, тем более близкими признаками обладают сравниваемые объекты.
Это называется _извлечение признаков_ - преобразование элемента в список чисел(координат), которые могут использоваться для сравнения.
Эта формула проста, но не является лучшей. На практике чаще применяют метрику близости косинусов, которая не измеряет расстояние между двумя векторами, а сравнивает между ними углы.
У алгоритма k-ближайших соседей есть два основных применения:
* классификация = распределение по категориям
* регресия      = прогнозирование ответа(в числовом выражении)

При работе с алгоритмом ближайших соседей, важно правильно выбрать признаки для сравнения. Под правильным выбором понимают:
* признаки должны быть напрямую связанны с объектом
* признаки не должны содержать смещения характеристик только в одну сторону, т.е. оценка не должна производиться только по одному признаку, а по многим

Эти алгоритмы применяются в машинном обучении, например при OCR - "Optical Character Recognition", оптическое распознавание текста.
Те же принципы и для распознавания изображений или речи, построения спам-фильтров и т.д. 
В основе любой сложной технологии лежит простая идея.
Первый шаг в таких алгоритмах называется - тренировкой.
Спам фильтры используют простой алгоритм называемый _наивным классификатором Байеса_



[Реализация алгоритма k ближайших соседей](chapter_10/k_nearest_neighbors.go)

[Упражнения](chapter_10/nearestneighbors_tasks.md)

## Глава 11. Что дальше?

# Бинарное дерево
Для каждого узла все узлы левого поддерева содержат _меньшие_ значения, а все узлы правого поддерева - _большие_ значения.
Поиск элемента в бинарном дереве в _среднем_ выполняется за время О(log n), а в _худшем случае_ - за время О(n). Несмотря на то, что отсортированный массив даже в _худшем случае_ выполняется за время O(log n) бинарное дерево в среднем работает быстрее при удалении и вставке элементов.
У деревьев есть свои недостатки:
* они не поддерживают произвольный доступ к i-му элементу
* среднее время операций зависит от сбалансированности дерева
Существуют специальные красно-черные деревья способные к самобалансировке.
В-деревья используются для хранения информации в базах данных.

# Инвертированные индексы
Используется в поисковых системах. 
Упрощенно: создается хэш-таблица где ключ -это контент, а значение адрес сайта. При поиске по контенту будет выдавать адрес. Получается как бы инвертированная система. Обычно ведь в картах лежит наоборот ключ - это адрес, а значение - это контент.

# Преобразование Фурье
Очень часто встречающийся алгоритм, которые показывает насколько велик каждая состовляющая из которых состоит элемент.
Используется в обработке сигналов, сжатия музыки (в частности мр3), графических форматах (jpg), для прогнозирования землятресений и анализа ДНК и во многих других.

# Параллельные алгоритмы

# MapReduce
Один из представитеелй распределенных алгоритмов.
Эти алгоритмы используются в тех ситуациях когда нужно выполнить большой объем работы за короткое время. 
В основе MapReduce лежат две простые идеи:
* функция отображения _map_
* функция свертки _reduce_

Функция map получает массив и применяет одну функцию к каждому элементу массива.
Функция reduce "сокращает" список элементов до одного, например суммируя их.

# Фильтры Блума и HyperLogLog

[Реализация алгоритма k ближайших соседей](chapter_11/)

[Упражнения](chapter_11/)











