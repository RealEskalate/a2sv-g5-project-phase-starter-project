import 'package:flutter/material.dart';

import '../../../authentication/domain/entity/user.dart';
import '../../domain/entity/product.dart';

import '../pages/detail_page.dart';

class ProductCard extends StatelessWidget {
  const ProductCard({super.key, required this.product, required this.user});
  final User user;
  final Product product;

  @override
  Widget build(BuildContext context) {
    return  GestureDetector(
            onTap: () {
              print('tapped');
              Navigator.of(context).push(
                    MaterialPageRoute(
                      builder: (BuildContext context) {
                        return DetailPage(product: product, user: user);
                      },
                    ),
                  );
            },  
                
            child: Card(
                child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                ClipRRect(
              borderRadius: const BorderRadius.only(
                topLeft: Radius.circular(8.0),
                topRight: Radius.circular(8.0),
              ),
              child: Image.network(
                product.image,
                fit: BoxFit.cover,
              ),
            ),
                
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(product.name),
                      Text("\$${product.price.toString()}"),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 8.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        product.seller.username,
                        style: const TextStyle(color: Colors.grey),
                      ),
                      Padding(
                        padding: const EdgeInsets.all(8.0),
                        child: const Row(
                          children: [
                            Icon(
                              Icons.star,
                              color: Colors.yellow,
                            ),
                            Text(
                              "(5)",
                              style: TextStyle(color: Colors.grey),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            )),
          );
  }
}

class TextFieldCard extends StatelessWidget {
  final title;
  final dollar;
  final fill;
  final area;
  TextEditingController text = TextEditingController();
  
  TextFieldCard(this.title,
      {required this.text,
      this.dollar = false,
      this.area = false,
      super.key,
      this.fill = true,
      });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(title),
        Stack(children: [
          TextField(
            controller: text,
            keyboardType: dollar ? TextInputType.number : TextInputType.text,
            maxLines: area ? 4 : 1,
            decoration: InputDecoration(
                filled: fill ? true : false,
                fillColor: const Color(0xFFF0F0F0),
                border:
                    OutlineInputBorder(borderRadius: BorderRadius.circular(5)),
                enabledBorder: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(5),
                    borderSide: BorderSide(
                        color: (fill) ? Colors.transparent : Colors.grey,
                        width: fill ? 2.0 : 1.0)),
                focusedBorder: const OutlineInputBorder(
                    borderSide: BorderSide(
                        color: Color.fromARGB(255, 83, 122, 249), width: 2))),
          ),
          dollar
              ? Positioned(
                  right: 0,
                  child: Container(
                    margin: const EdgeInsets.only(
                        left: 5, right: 5, bottom: 5, top: 5),
                    color: const Color(0xFFF3F3F3),
                    padding: const EdgeInsets.all(15),
                    child: const Align(
                      alignment: Alignment.bottomRight,
                      child: Text('\$'),
                    ),
                  ),
                )
              : Container(),
        ])
      ],
    );
  }
}
