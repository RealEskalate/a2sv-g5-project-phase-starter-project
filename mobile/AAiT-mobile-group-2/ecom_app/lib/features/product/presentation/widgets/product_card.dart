import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';

import '../../domain/entities/product.dart';
// import '../models/shoe.dart';

class ProductCard extends StatelessWidget {
  final Product product;
  const ProductCard({
    super.key,
    required this.product,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
        clipBehavior: Clip.hardEdge,
        child: InkWell(
          onTap: () {
            Navigator.pushNamed(context, '/details',
                arguments: Product(
                    id: product.id,
                    imageUrl: product.imageUrl,
                    name: product.name,
                    price: product.price,
                    description: product.description));
          },
          splashColor: Colors.blue[100],
          child: SizedBox(
            width: 380,
            height: 280,
            child: Column(
              children: [
                Flexible(
                  flex: 2,
                  child: SizedBox(
                      width: double.infinity,
                      child: Opacity(
                        opacity: 0.8,
                        child: CachedNetworkImage(
                          imageUrl: product.imageUrl,
                          placeholder: (context, url) => const SizedBox(
                              height: 185,
                              child: const CircularProgressIndicator()),
                          errorWidget: (context, url, error) =>
                              const Icon(Icons.error),
                          fit: BoxFit.cover,
                        ),
                      )),
                ),
                Flexible(
                    flex: 1,
                    child: Padding(
                      padding: const EdgeInsets.all(10),
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.spaceAround,
                        children: [
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              Flexible(
                                child: Text(
                                  product.name,
                                  style: const TextStyle(
                                      fontWeight: FontWeight.w600,
                                      fontSize: 20),
                                  overflow:
                                      TextOverflow.ellipsis, // Handle overflow
                                ),
                              ),
                              Text(
                                '\$${product.price}',
                                style: const TextStyle(
                                    fontWeight: FontWeight.w500),
                              )
                            ],
                          ),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              const Flexible(
                                child: Text(
                                  "Chuck Taylor's",
                                  style: const TextStyle(color: Colors.black45),
                                  overflow:
                                      TextOverflow.ellipsis, // Handle overflow
                                ),
                              ),
                              Row(
                                children: [
                                  Icon(
                                    Icons.star_rate,
                                    color: Colors.amber[400],
                                  ),
                                  const Text(
                                    '(4.0)',
                                    style:
                                        const TextStyle(color: Colors.black45),
                                  )
                                ],
                              )
                            ],
                          )
                        ],
                      ),
                    ))
              ],
            ),
          ),
        ));
  }
}
