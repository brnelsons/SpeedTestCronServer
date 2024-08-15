export class Service {
    getHistory() {
        return fetch("api/v1/history")
            .then(response => {
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                return response.json();
            })
    }
}