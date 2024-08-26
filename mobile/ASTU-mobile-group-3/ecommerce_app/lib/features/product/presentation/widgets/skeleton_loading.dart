import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';
import 'package:shimmer/shimmer.dart';

class SkeletonLoading extends StatelessWidget {
  const SkeletonLoading({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          height: 160,
          margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
          decoration: const BoxDecoration(
            color: MyTheme.skeletonColor2,
            borderRadius: BorderRadius.only(
              topLeft: Radius.circular(20),
              topRight: Radius.circular(20),
            ),
          ),
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Container(
              height: 20,
              width: MediaQuery.of(context).size.width / 2,
              margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 0),
              decoration: const BoxDecoration(
                color: MyTheme.ecInputGrey,
                borderRadius: BorderRadius.all(Radius.circular(5)),
              ),
            ),
            Container(
              height: 20,
              width: MediaQuery.of(context).size.width / 5,
              margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 0),
              decoration: const BoxDecoration(
                color: MyTheme.skeletonColor1,
                borderRadius: BorderRadius.all(Radius.circular(5)),
              ),
            ),
          ],
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Container(
              height: 20,
              width: MediaQuery.of(context).size.width / 4,
              margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 0),
              decoration: const BoxDecoration(
                color: MyTheme.skeletonColor1,
                borderRadius: BorderRadius.all(Radius.circular(5)),
              ),
            ),
            Container(
              height: 20,
              width: MediaQuery.of(context).size.width / 3,
              margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
              decoration: const BoxDecoration(
                color: MyTheme.ecInputGrey,
                borderRadius: BorderRadius.all(Radius.circular(3)),
              ),
            ),
          ],
        )
      ],
    );
  }
}

class ShimmerList extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 8),
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: Colors.white54,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Shimmer.fromColors(
            baseColor: Colors.grey.shade300,
            highlightColor: Colors.grey.shade100,
            child: ClipRRect(
              borderRadius:
                  const BorderRadius.vertical(top: Radius.circular(10)),
              child: Container(
                height: 150,
                color: Colors.grey.shade300,
              ),
            ),
          ),
          const SizedBox(height: 16),
          Shimmer.fromColors(
            baseColor: Colors.grey.shade300,
            highlightColor: Colors.grey.shade100,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Container(
                height: 20,
                width: 200,
                color: Colors.grey.shade300,
              ),
            ),
          ),
          const SizedBox(height: 8),
          Shimmer.fromColors(
            baseColor: Colors.grey.shade300,
            highlightColor: Colors.grey.shade100,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Container(
                height: 20,
                width: 200,
                color: Colors.grey.shade300,
              ),
            ),
          ),
          const SizedBox(height: 16),
          Shimmer.fromColors(
            baseColor: Colors.grey.shade300,
            highlightColor: Colors.grey.shade100,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Container(
                height: 14,
                width: 150,
                color: Colors.grey.shade300,
              ),
            ),
          ),
          const SizedBox(height: 8),
          Shimmer.fromColors(
            baseColor: Colors.grey.shade300,
            highlightColor: Colors.grey.shade100,
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Row(
                children: [
                  const Icon(Icons.star, color: Colors.grey, size: 15),
                  const SizedBox(width: 8),
                  Container(
                    height: 14,
                    width: 30,
                    color: Colors.grey.shade300,
                  ),
                ],
              ),
            ),
          ),
          const SizedBox(height: 16),
        ],
      ),
    );
  }
}
