import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class ProductCard extends StatelessWidget {
  final String imageUrl;
  final String productName;
  final String productType;
  final int price;
  final String rating;

  const ProductCard({
    super.key,
    required this.imageUrl,
    required this.price,
    required this.productName,
    required this.productType,
    required this.rating,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
      decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(20),
          boxShadow: const [
            BoxShadow(
                color: MyTheme.shadowColor, spreadRadius: 2, blurRadius: 10)
          ]),
      child: Column(
        children: [
          ClipRRect(
            borderRadius: const BorderRadius.only(
                topLeft: Radius.circular(20), topRight: Radius.circular(20)),
            child: SizedBox(
              height: 160,
              width: double.infinity,
              child: Image.network(
                imageUrl,
                loadingBuilder: (
                  context,
                  thisChild,
                  progress,
                ) {
                  if (progress == null) {
                    return thisChild;
                  } else {
                    return Container(
                      height: 160,
                      width: double.infinity,
                      decoration: const BoxDecoration(
                        color: Color.fromARGB(255, 207, 207, 207),
                      ),
                      child: const Stack(children: [
                        Positioned(
                          bottom: 10,
                          right: 10,
                          child: CircularProgressIndicator(),
                        ),
                      ]),
                    );
                  }
                },
                errorBuilder: (BuildContext context, Object exception,
                    StackTrace? stackTrace) {
                  return Container(
                    height: 160,
                    width: double.infinity,
                    decoration: const BoxDecoration(
                      color: Color.fromARGB(255, 207, 207, 207),
                    ),
                    child: const Stack(children: [
                      Positioned(
                        bottom: 10,
                        right: 10,
                        child: CircularProgressIndicator(),
                      ),
                    ]),
                  );
                },
                fit: BoxFit.fitWidth,
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  productName,
                  style: const TextStyle(
                      fontFamily: 'poppins',
                      fontSize: 20,
                      fontWeight: FontWeight.bold),
                ),
                Text(
                  '$price\$',
                  style: const TextStyle(
                    fontFamily: 'poppins',
                    fontWeight: FontWeight.bold,
                    fontSize: 12,
                  ),
                ),
              ],
            ),
          ),
          Padding(
            padding: const EdgeInsets.only(left: 20, right: 20, bottom: 10),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                ConstrainedBox(
                  constraints: BoxConstraints(
                      maxWidth: MediaQuery.of(context).size.width - 200),
                  child: Text(
                    productType,
                    style: const TextStyle(
                      color: MyTheme.ecTextGrey,
                      fontSize: 14,
                      fontFamily: 'poppins',
                      overflow: TextOverflow.ellipsis,
                    ),
                  ),
                ),
                Wrap(
                  children: [
                    const Icon(
                      Icons.star,
                      color: Colors.yellow,
                    ),
                    Text(
                      '($rating)',
                      style: const TextStyle(
                        color: MyTheme.ecTextGrey,
                        fontSize: 14,
                        fontFamily: 'poppins',
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
