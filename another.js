const isPrime = (n) => {
    if (n < 2) {
        return false;
    }

    const max = Math.sqrt(n);

    for (let i = 2; i <= Number.parseInt(max); i++) {
        if (n % i === 0) {
            return false;
        }
    }

    return true;
};

const run = () => {
    let sum = 0
    for (i = 0; i < 2_000_000; i++) {
        if (isPrime(i)) {
            sum += i
        }
    }

    console.log(sum)
};

run();
