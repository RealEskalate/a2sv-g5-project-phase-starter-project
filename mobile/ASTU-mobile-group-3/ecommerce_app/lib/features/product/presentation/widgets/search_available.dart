import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';
import '../pages/search_product_page.dart';

class SearchNavigator extends StatelessWidget {
  const SearchNavigator({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(
        horizontal: 30,
        vertical: 10,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          const Text(
            'Available Product',
            style: TextStyle(fontWeight: FontWeight.bold, fontSize: 30),
          ),
          Container(
            decoration: BoxDecoration(
                border: Border.all(color: MyTheme.ecGrey, width: 2),
                borderRadius: BorderRadius.circular(16)),
            child: IconButton(
              onPressed: () {
                Navigator.pushNamed(context, SearchProduct.routes);
              },
              icon: const Icon(
                Icons.search,
                color: MyTheme.ecGrey,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
