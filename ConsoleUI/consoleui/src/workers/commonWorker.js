export async function HandlerWrapper(name, handler) {
    try {
        console.log(`resting ${name}`)
        return await handler();
    }
    catch (error) {
        console.log(error.error, error.config, error.code, error.request, error.response, error.response?.data);
    }
}

export function GetJsonHeader() {
    return {
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        }
    }
}

export function GetTokenHeader(token) {
    return  {
        headers: {
        "Accept": "application/json",
        "Authorization": "Bearer " + token  // передача токена в заголовке
        }
    };
}


