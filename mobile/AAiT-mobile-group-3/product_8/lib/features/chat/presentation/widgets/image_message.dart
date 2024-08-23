import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';


class ImageMessage extends StatelessWidget {
  const ImageMessage({super.key});

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.only(left: 15),
      shadowColor: Colors.white,
      clipBehavior: Clip.antiAlias,
      child: CachedNetworkImage(
        imageUrl:
            'https://bestinau.com.au/wp-content/uploads/2019/01/free-images.jpg',
      ),
    );
  }
}
