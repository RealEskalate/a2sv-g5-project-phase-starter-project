import 'package:flutter/material.dart';

import '../../../auth/presentation/widgets/custom_text.dart';
import '../../domain/entities/product_entity.dart';
import 'image_loader.dart';

class ProductListItem extends StatelessWidget {
  final ProductEntity product;

  const ProductListItem({
    super.key,
    required this.product,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 10,
      child: Container(
        width: 366,
        height: 220,
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.only(
            topLeft: Radius.circular(16),
            topRight: Radius.circular(16),
          ),
        ),
        child: Column(
          children: [
            imageLoader(product.imageUrl),
            const Spacer(),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16),
              child: Row(
                children: [
                  Text(
                    product.name,
                    style: const TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  const Spacer(),
                  Text(
                    '\$${product.price}',
                    style: const TextStyle(
                      fontSize: 14,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                ],
              ),
            ),
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 16),
              child: Row(
                children: [
                  CustomText(
                    text: 'Category',
                    color: Colors.grey,
                  ),
                  Spacer(),
                  Icon(
                    Icons.star,
                    color: Colors.yellow,
                  ),
                  CustomText(
                    text: '(4.0)',
                    color: Colors.grey,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
