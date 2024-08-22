import 'package:flutter/material.dart';
import '../../domain/entities/product.dart';
import 'item_card.dart';
class ProductList extends StatelessWidget {
  final List<Product> products;

  const ProductList({super.key, required this.products});

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemCount: products.length,
      itemBuilder: (context, index) {
        final product = products[index];
        return GestureDetector(
          onTap: () {
            Navigator.pushNamed(
              context,
              '/details',
              arguments: product,
            );
          },


          child: ProductItemCard(product: product),
          

        );
      },
    );
  }
}