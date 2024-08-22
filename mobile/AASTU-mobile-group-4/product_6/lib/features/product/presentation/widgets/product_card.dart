import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../domain/entities/product.dart';
import '../pages/details_page.dart';

class ProductCard extends StatelessWidget {
  final Product productObject;
  const ProductCard({super.key, required this.productObject});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) => DetailsPage(
              productObject: productObject,
            ),
          ),
        );
      },
      child: Container(
        margin: const EdgeInsets.only(left: 10, right: 10, top: 25),
        height: 220,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(13),
          color: Colors.white,
          shape: BoxShape.rectangle,
          boxShadow: [
            BoxShadow(
              color: Colors.grey.withOpacity(0.5),
              spreadRadius: 1,
              blurRadius: 1,
              offset: const Offset(3, 5),
            ),
          ],
        ),
        // color: Colors.red,
        child: Column(
          mainAxisSize: MainAxisSize.max,
          children: [
            SizedBox(
                height: 150,
                child: Image.network(
                  productObject.imageUrl,
                )),
            Container(
              padding: const EdgeInsets.all(8),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    productObject.name,
                    style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                        fontWeight: FontWeight.w600,
                        fontSize: 16,
                        color: Color.fromRGBO(62, 62, 62, 1),
                      ),
                    ),
                  ),
                  Text(
                    productObject.price.toString(),
                    style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 14.0,
                    )),
                  )
                ],
              ),
            ),
            // const SizedBox(height: 2),
            Padding(
              padding: const EdgeInsets.only(left: 8, right: 8),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    'Men',
                    style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                        fontWeight: FontWeight.w600,
                        color: Color.fromRGBO(170, 170, 170, 1),
                        fontSize: 12.0,
                        height: 0.5,
                      ),
                    ),
                  ),
                  Row(
                    // mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      Container(
                        padding: const EdgeInsets.all(0),
                        height: 24,
                        width: 24,
                        child: Center(
                          child: IconButton(
                            onPressed: () {},
                            icon: const Icon(
                              Icons.star,
                              size: 12,
                              color: Color.fromRGBO(255, 215, 0, 1),
                            ),
                          ),
                        ),
                      ),
                      Text(
                        // ignore: prefer_adjacent_string_concatenation
                        '(4.0)',
                        style: GoogleFonts.poppins(
                            textStyle: const TextStyle(
                          fontWeight: FontWeight.w400,
                          fontSize: 12.0,
                          color: Color.fromRGBO(170, 170, 170, 1),
                        )),
                      ),
                    ],
                  )
                ],
              ),
            )
          ],
        ),
      ),
    );
  }
}
