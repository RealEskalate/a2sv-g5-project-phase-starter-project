import 'package:flutter/material.dart';

import '../../../../../../core/const/width_height.dart';

class SizeOfProduct extends StatelessWidget {
  const SizeOfProduct({super.key});

  @override
  Widget build(BuildContext context) {
    List<int> size = [39, 40, 41, 42, 43, 44, 45];
    double width = WidthHeight.screenWidth(context);
    double height = WidthHeight.screenHeight(context);
    return SizedBox(
      height: 60, // Fixed height for the ListView
      child: ListView.builder(
        itemCount: size.length,
        scrollDirection: Axis.horizontal,
        itemBuilder: (context, index) {
          int current = size[index];
          return Container(
            alignment: Alignment.center,
            margin: const EdgeInsets.all(5),
            width: width*0.116,
            height: height*0.0626,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
              color: current != 41 ? Colors.white : Colors.blue,
              boxShadow: const [
                BoxShadow(
                  color: Color.fromARGB(255, 223, 222, 222),
                  offset: Offset(0, 1),
                  spreadRadius: 0.4,
                )
              ],
            ),
            child: Text(
              '$current',
              style: TextStyle(
                color: current != 41
                    ? const Color.fromARGB(255, 3, 3, 3)
                    : const Color.fromARGB(255, 255, 255, 255),
                fontSize: 20,
              ),
            ),
          );
        },
      ),
    );
  }
}
