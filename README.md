### LRU (Least Recently Used) Cache

**LRU** is a caching algorithm that evicts the least recently accessed items when the cache reaches its capacity.

This strategy is useful in scenarios where recently accessed data is likely to be used again soon — such as caching database queries or frequently requested images.

The **TTL (Time To Live)** modification adds expiration logic to each item in the cache. Once an item’s TTL runs out, it gets automatically removed. This prevents stale data from occupying space and further improves cache efficiency.

---

### LRU (Least Recently Used) Cache

**LRU** — это алгоритм кэширования, при котором удаляются элементы, к которым давно не обращались, когда кэш переполняется.

Такой механизм эффективен в ситуациях, где есть высокая вероятность повторного использования недавно запрошенных данных: например, при кэшировании запросов к базе данных или загрузке изображений.

Модификация с **TTL (Time To Live)** позволяет автоматически удалять элементы после истечения их срока жизни. Это защищает кэш от "протухших", "неактуальных" данных и делает его ещё более полезным.
