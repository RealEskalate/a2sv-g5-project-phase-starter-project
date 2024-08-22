import 'package:flutter/material.dart';

Widget buildSearchButton(BuildContext context) {
  return Padding(
    padding: const EdgeInsets.all(8.0),
    child: Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        border: Border.all(
          color: const Color.fromRGBO(217, 217, 217, 1),
          width: 1.0,
        ),
      ),
      child: Tooltip(
        message: 'Search Products',
        child: Material(
          color: Colors.transparent,
          child: InkWell(
            borderRadius: BorderRadius.circular(10),
            onTap: () {
              Navigator.pushNamed(context, '/search');
            },
            child: const Padding(
              padding: EdgeInsets.all(8.0),
              child: Icon(
                Icons.search,
                size: 30,
                color: Color.fromRGBO(217, 217, 217, 1),
              ),
            ),
          ),
        ),
      ),
    ),
  );
}
