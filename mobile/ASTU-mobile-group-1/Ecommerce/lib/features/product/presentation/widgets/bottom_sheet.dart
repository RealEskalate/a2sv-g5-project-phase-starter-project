import 'package:flutter/material.dart';

Future<dynamic> bottomSheet(BuildContext context) {
  RangeValues _priceRange = const RangeValues(20, 80);
  return showModalBottomSheet(
    context: context,
    builder: (BuildContext context) {
      RangeValues tempPriceRange = _priceRange;
      return StatefulBuilder(
        builder: (BuildContext context, StateSetter setState) {
          return Padding(
            padding: const EdgeInsets.all(32.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisSize: MainAxisSize.min,
              children: [
                const TextField(
                  decoration: InputDecoration(
                    labelText: 'Category',
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 16),
                const Text(
                  'Price',
                  style: TextStyle(fontSize: 16),
                ),
                RangeSlider(
                  values: tempPriceRange,
                  min: 0,
                  max: 100,
                  divisions: 10,
                  activeColor: Colors.blue,
                  inactiveColor: Colors.grey,
                  onChanged: (RangeValues values) {
                    setState(() {
                      tempPriceRange = values;
                    });
                  },
                ),
                const SizedBox(height: 16),
                ElevatedButton(
                  onPressed: () {
                    setState(() {
                      _priceRange = tempPriceRange;
                    });
                    Navigator.of(context).pop();
                  },
                  style: ElevatedButton.styleFrom(
                    minimumSize: const Size(
                        double.infinity, 48), // Make button full width
                    backgroundColor: const Color(0xff3F51F3),
                    shape: RoundedRectangleBorder(
                      borderRadius:
                          BorderRadius.circular(8), // Background color
                    ),
                  ),
                  child: const Text(
                    'APPLY',
                    style: TextStyle(fontSize: 16, color: Color(0xffFFFFFF)),
                  ),
                ),
                const SizedBox(
                  height: 10,
                ),
              ],
            ),
          );
        },
      );
    },
  );
}
