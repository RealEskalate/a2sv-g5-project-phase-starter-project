import 'package:flutter/material.dart';
import 'package:shimmer/shimmer.dart';

class ListLoadingShimmer extends StatelessWidget {
  const ListLoadingShimmer({super.key});

  @override
  Widget build(BuildContext context) {
    return Shimmer.fromColors(
        key: const Key('LIST_SHIMMER'),
        baseColor: Colors.grey.shade300,
        highlightColor: Colors.grey.shade100,
        enabled: true,
        child: SizedBox(
          height: MediaQuery.of(context).size.height - 150,
          child: const SingleChildScrollView(
            scrollDirection: Axis.vertical,
            padding: EdgeInsets.all(16.0),
            child: Column(
              children: [
                CardPlaceHolder(),
                CardPlaceHolder(),
                CardPlaceHolder(),
                CardPlaceHolder()
              ],
            ),
          ),
        ));
  }
}

class CardPlaceHolder extends StatelessWidget {
  const CardPlaceHolder({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(top: 5, bottom: 5),
      child: Card(
        elevation: 5,
        child: Container(
          width: double.infinity,
          height: 200,
          margin: const EdgeInsets.all(16.0),
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(12.0),
            color: Colors.white,
          ),
        ),
      ),
    );
  }
}

class ImageLoadingShimmer extends StatelessWidget {
  const ImageLoadingShimmer({super.key});

  @override
  Widget build(BuildContext context) {
    return Shimmer.fromColors(
        key: const Key('IMAGE_SHIMMER'),
        baseColor: Colors.grey.shade300,
        highlightColor: Colors.grey.shade100,
        child: Container(
          width: 100,
          height: 100,
          color: Colors.white,
        ));
  }
}
