export const priceDataFormater = (number) => {
    if (number > 1000000000) {
        return (number / 1000000000).toString() + ' میلیارد ';
    } else if (number > 1000000) {
        return (number / 1000000).toString() + ' میلیون ';
    } else if (number > 1000) {
        return (number / 1000).toString() + ' هزار ';
    } else {
        return number.toString();
    }
}