import 'dart:io';

import 'package:flutter/material.dart';

import '../../../domain/entities/product_entity.dart';
import 'styles/text_style.dart';

class MyCardBox extends StatefulWidget {
  final ProductEntity product;
  const MyCardBox({
    super.key,
    required this.product,
  });

  @override
  State<MyCardBox> createState() => _MyCardBoxState();
}

class _MyCardBoxState extends State<MyCardBox> {
  @override
  Widget build(BuildContext context) {
  bool isFile = File(widget.product.imageUrl).existsSync();
    return GestureDetector(
      onTap: () {
       Navigator.pushNamed(context, '/details_page',arguments: widget.product);
      },
      child: Card(
        child: Container(
          alignment: Alignment.center,
          width: MediaQuery.of(context).size.width,
          height: 280,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(16.0),
          ),
          child: Column(
            children: [
              SizedBox(
                width: MediaQuery.of(context).size.width,
                height: 200,
                child: ClipRRect(
                    borderRadius: const BorderRadius.only(
                        topLeft: Radius.circular(16.0),
                        topRight: Radius.circular(16.0)),
                    child: isFile ? Image.file(File(widget.product.imageUrl)) : Image.network(widget.product.imageUrl, fit: BoxFit.cover)),
              ),
              Container(
                padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 12.0),
                child: Column(
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        CustomTextStyle(
                            name: widget.product.name,
                            weight: FontWeight.w500,
                            size: 20.0),
                        CustomTextStyle(
                            name: '\$${widget.product.price}', weight: FontWeight.w500, size: 14),
                      ],
                    ),
                    const Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        CustomTextStyle(
                          name: "Men's shoe",
                          weight: FontWeight.w400,
                          size: 12,
                          color: Color.fromRGBO(170, 170, 170, 1.0),
                        ),
                        Row(
                          children: [
                            Icon(
                              Icons.star,
                              color: Color.fromRGBO(255, 215, 0, 1),
                            ),
                            CustomTextStyle(
                                name: '(4)',
                                weight: FontWeight.w400,
                                size: 12,
                                color: Color.fromRGBO(170, 170, 170, 1.0),
                                family: 'Sora'),
                          ],
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
