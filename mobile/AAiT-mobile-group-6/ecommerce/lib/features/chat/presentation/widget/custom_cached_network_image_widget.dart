import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';

import '../../../product/domain/entitity/product.dart';

class CustomCachedNetworkImage extends StatelessWidget {
  const CustomCachedNetworkImage({
    super.key,
    required this.product,
    this.height = 400,
  });

  final Product product;
  final double height;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: height,
      width: double.infinity,
      child: ClipRect(
        child: FittedBox(
          fit: BoxFit.cover,
          child: CachedNetworkImgeWithshimmerWaiter(imageUrl: product.imageUrl),
        ),
      ),
    );
  }
}

class CachedNetworkImgeWithshimmerWaiter extends StatelessWidget {
  const CachedNetworkImgeWithshimmerWaiter({
    super.key,
    required this.imageUrl,
  });

  final String imageUrl;

  @override
  Widget build(BuildContext context) {
    return CachedNetworkImage(
      imageUrl: imageUrl,
      placeholder: (context, url) => const CircularProgressIndicator(),
      errorWidget: (context, url, error) => const Icon(Icons.error),
    );
  }
}
