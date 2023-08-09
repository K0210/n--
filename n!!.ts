function f(x: number): number{
    let y: number = 1
    if(x > 0){
        let k: number = x
        while(k >= 1){
            y *= k
            k -= 2
        }
    }
    else if(x <= 0 && x %2 != 0){
        let k: number = 1
        while(k >= x){
            y = y / k
            k = k - 2
        }
    }
    else if(x != 0){
        alert("error")
        y = Number.NaN
    }
    return y
}
