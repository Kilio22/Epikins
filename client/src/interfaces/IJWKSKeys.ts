export interface IJWKSKey {
    kid: string,
    x5c: string[]
}

export interface IJWKSKeys {
    keys: IJWKSKey[]
}