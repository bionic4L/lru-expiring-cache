/*
Данный код создан для того, чтобы Вы могли сами посмотреть как работает
LRU кэш. Вы можете установить свои лимиты, типы данных, хранимых в кэше.

Код представляет из себя "песочницу", чтобы новички гоферы могли
на практике посмотреть как работает этот механизм.

Это не отменяет факта, что лучше самому реализовать кеш с нуля для
полного понимания.
*/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	lruCache "lru-and-lfu-cache/lru-cache"
	"os"
	"strings"
	"time"
)

const (
	CacheCapacity = 10
	TTL           = 10 * time.Second
	CronInterval  = "@every 60s"
)

func Printer(cache *lruCache.LRU) {
	fmt.Println("\ndoubly linked list (from new to old):")
	for e := cache.Queue.Front(); e != nil; e = e.Next() {
		item := e.Value.(*lruCache.Item)
		fmt.Printf("key: %v, value: %v\n", item.Key, item.Value)
	}

	fmt.Println("\nmap:")
	cache.Items.Range(func(k, v any) bool {
		elem := v.(*list.Element)
		item := elem.Value.(*lruCache.Item)
		fmt.Printf("key: %v, value: %v\n", item.Key, item.Value)
		return true
	})
	fmt.Println()
}

func main() {
	cache := lruCache.NewLRU(CacheCapacity, TTL, CronInterval)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter one of the following commands:")
	fmt.Println("- to insert/update objects 'set,<key>,<value>' ")
	fmt.Println("- to get values 'get,<key>' ")
	fmt.Println("- press enter with empty line to exit.")

	for {
		fmt.Print("-> ")

		if !scanner.Scan() {
			fmt.Println("scan error")
			break
		}
		input := scanner.Text()

		if input == "" {
			fmt.Println("exiting...")
			break
		}

		parts := strings.SplitN(input, ",", 3)
		if len(parts) < 2 {
			fmt.Println("invalid input, use 'set,<key>,<value>' or 'get,<key>'")
			continue
		}

		command := strings.ToLower(parts[0])
		key := parts[1]

		switch command {
		case "set":
			if len(parts) != 3 {
				fmt.Println("invalid set format, use 'set,<key>,<value>'")
				continue
			}
			value := parts[2]
			cache.Set(key, value)
			fmt.Printf("set key: %v, value: %v\n", key, value)
			Printer(cache)

		case "get":
			val := cache.Get(key)
			if val == nil {
				fmt.Printf("key '%s' not found.\n", key)
			} else {
				fmt.Printf("got value: %v\n", val)
			}
			Printer(cache)

		default:
			fmt.Println("unknown command, use 'set' or 'get'")
		}
	}
}
