// implement class so that i dont have to make boilerplate objects
class Store<T> {
    private value = $state() as T
    constructor(initialValue: T) {
        this.value = initialValue
    }

    // getter and setter to update
    get(): T {
        return this.value
    }
    set(newValue: T) {
        this.value = newValue
    }
}
const createStore = <T>(initialValue: T) => new Store<T>(initialValue)
export const isSidebarToggled = createStore(false)