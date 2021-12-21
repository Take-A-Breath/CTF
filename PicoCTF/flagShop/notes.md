# Flag Shop
## Category: General Skills

#### Description:
> There's a flag shop selling stuff, can you buy a flag? Source. Connect with nc jupiter.challenges.picoctf.org 4906.
 * Source file provided: `store.c`


```Console
 Enter a menu selection
2
Currently for sale
1. Defintely not the flag Flag
2. 1337 Flag
1
These knockoff Flags cost 900 each, enter desired quantity
300000000

The final cost is: -582939648

Your current balance after transaction: 582940748

Welcome to the flag exchange
We sell flags

1. Check Account Balance

2. Buy Flags

3. Exit

 Enter a menu selection
2
Currently for sale
1. Defintely not the flag Flag
2. 1337 Flag
2
1337 flags cost 100000 dollars, and we only have 1 in stock
Enter 1 to buy one1
YOUR FLAG IS: picoCTF{m0n3y_bag5_9c5fac9b}
```
 * exploit:
```Console
if(auction_choice == 1){
    printf("These knockoff Flags cost 900 each, enter desired quantity\n");
    int number_flags = 0;
    fflush(stdin);
    scanf("%d", &number_flags);
    if(number_flags > 0){
        int total_cost = 0;
        total_cost = 900*number_flags;
        printf("\nThe final cost is: %d\n", total_cost);
        if(total_cost <= account_balance){
            account_balance = account_balance - total_cost;
            printf("\nYour current balance after transaction: %d\n\n", account_balance);
        }
        else{
            printf("Not enough funds to complete purchase\n");
        }
    }
}
```

> Notice that total_cost is declared as an int. This is specifically a Signed Integer, which has a range between -2,147,483,648 to
> 2,147,483,647. Signed integers use the first bit to store whether it is negative or positive, 0 indicating positive and 1 indicating
> negative. What happens if you add 1 to 2,147,483,647 and store the result in a signed integer? Well the first bit goes from 0 to 1, meaning
> that the number is now negative! In fact, due to the way Two's Complement, the method used to represent negative numbers in binary, works,
> it actually wraps around to the most negative integer: -2,147,483,648.

__flag:__ picoCTF{m0n3y_bag5_9c5fac9b}
