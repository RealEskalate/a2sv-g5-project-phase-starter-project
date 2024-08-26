import 'package:flutter/material.dart';

import '../../../../authentication/domain/entities/user_data.dart';
import 'styles/text_style.dart';

class Seller extends StatelessWidget {
  final UserEntity seller;
  const Seller({super.key, required this.seller});

  @override
  Widget build(BuildContext context) {
    return Card(
      color: const Color.fromARGB(216, 255, 255, 255),
      elevation: 3,
      child: Padding(
        padding: const EdgeInsets.symmetric(vertical: 5.0, horizontal: 4),
        child: Row(
          mainAxisSize: MainAxisSize.max,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                CustomTextStyle(
                  name: seller.name,
                  size: 18,
                  weight: FontWeight.bold,
                ),
                const SizedBox(height: 5,),
                CustomTextStyle(
                  name: seller.email,
                  size: 14,
                  weight: FontWeight.w400,
                  color: const Color.fromARGB(225, 170, 170, 170),
                ),
              ],
            ),
            FilledButton.icon(
              icon: const Icon(Icons.send,),
              style: ButtonStyle(
                backgroundColor: WidgetStatePropertyAll(Theme.of(context).primaryColor),
                shape: WidgetStatePropertyAll(RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(3),
                )),
              ),
              onPressed: () {},
              label: const CustomTextStyle(
                  name: 'Contact Seller',
                  size: 16,
                  weight: FontWeight.normal,
                  color: Colors.white,
                )
            ),
          ],
        ),
      ),
    );
  }
}