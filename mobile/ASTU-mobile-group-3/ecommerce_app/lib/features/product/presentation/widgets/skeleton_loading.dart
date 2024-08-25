import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

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
