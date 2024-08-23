import 'package:flutter/material.dart';

import '../../domain/entities/product_entity.dart';
import '../pages/product_details_page.dart';

class ProductCard extends StatelessWidget {
  final Product product;
  const ProductCard({super.key, required this.product});

  @override
  Widget build(BuildContext context) {
    return Card(
      clipBehavior: Clip.hardEdge,
      child: InkWell(
        onTap: () => {
          Navigator.pushNamed(context, Detailspage.routeName,
              arguments: Product(
                  name: product.name,
                  imageUrl: product.imageUrl,
                  price: product.price,
                  description: product.description,
                  id: product.id))
        },
        splashColor: Colors.indigoAccent.shade400,
        child: Column(
          children: [
            Flexible(
                child: FractionallySizedBox(
                    widthFactor: 1.0,
                    heightFactor: 0.95,
                    child: ClipRRect(
                      borderRadius: const BorderRadius.only(
                          topRight: Radius.circular(20),
                          topLeft: Radius.circular(20)),
                      child: Image.network(
                        product.imageUrl,
                        fit: BoxFit.cover,
                      ),
                    ))),
            const SizedBox(
              height: 10.0,
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    product.name,
                    style: const TextStyle(
                        fontSize: 20.0, fontWeight: FontWeight.bold),
                  ),
                  Text(
                    '\$${product.price}',
                    style: const TextStyle(
                        fontSize: 15.0, fontWeight: FontWeight.bold),
                  )
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(left: 10.0, right: 10.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  const Text(
                    "Men's shoe",
                    style: TextStyle(fontSize: 10.0, color: Colors.black45),
                  ),
                  Row(
                    children: [
                      Icon(
                        Icons.star,
                        color: Colors.yellow[700],
                      ),
                      const Text(
                        '(4.0)',
                        style: TextStyle(fontSize: 10.0, color: Colors.black45),
                      ),
                    ],
                  )
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
