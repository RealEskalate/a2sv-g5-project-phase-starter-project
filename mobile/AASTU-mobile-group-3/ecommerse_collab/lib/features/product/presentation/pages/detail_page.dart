import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../service_locator.dart';
import '../../domain/entity/product.dart';
import '../../domain/usecase/add_product.dart';
import '../../domain/usecase/delete_product.dart';
import '../../domain/usecase/get_all_product.dart';
import '../../domain/usecase/get_product.dart';
import '../../domain/usecase/update_product.dart';
import '../bloc/blocs.dart';
import '../bloc/events.dart';
import '../bloc/states.dart';

import 'update_page.dart';

class DetailPage extends StatefulWidget {
  const DetailPage({super.key, required this.product});
  final Product product;

  @override
  State<DetailPage> createState() => _DetailPageState();
}

class _DetailPageState extends State<DetailPage> {
  int current = 39;

  Widget sizeNumber(number) {
    return GestureDetector(
      onTap: () {
        setState(() {
          current = int.parse(number);
          debugPrint('$current');
        });
      },
      child: Container(
          width: 60,
          height: 60,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(8),
            boxShadow: [
              BoxShadow(
                  color: (current != int.parse(number))
                      ? const Color.fromARGB(255, 236, 236, 247)
                          .withOpacity(0.5)
                      : const Color.fromARGB(255, 7, 80, 214)),
            ],
          ),
          padding: const EdgeInsets.symmetric(horizontal: 8.0),
          margin: const EdgeInsets.symmetric(horizontal: 4.0, vertical: 8.0),
          child: Center(
            child: Text(number,
                style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.w500,
                    color: (current == int.parse(number))
                        ? Colors.white
                        : Colors.black)),
          )),
    );
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => ProductBloc(
          addProductUseCase: getIt<AddProductUseCase>(),
          deleteProductUseCase: getIt<DeleteProductUseCase>(),
          updateProductUseCase: getIt<UpdateProductUseCase>(),
          getAllProductUseCase: getIt<GetAllProductUseCase>(),
          getProductUseCase: getIt<GetProductUseCase>()),
      child: Scaffold(
        body: SingleChildScrollView(
          child: SizedBox(
            height: MediaQuery.of(context).size.height,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Stack(children: [
                  SizedBox(
                      width: 366,
                      height: 250,
                      child: Image.network(widget.product.image)),
                  Positioned(
                      top: 20,
                      left: 15,
                      child: Container(
                        width: 40,
                        height: 40,
                        decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(40),
                            color: Colors.white),
                        child: IconButton(
                          icon: const Icon(Icons.arrow_back_ios_new_rounded,
                              size: 18, color: Colors.blue),
                          onPressed: () {
                            Navigator.of(context).pop();
                          },
                        ),
                      ))
                ]),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        widget.product.category ?? '',
                        style: const TextStyle(color: Colors.grey),
                      ),
                      const Row(
                        children: [
                          Icon(
                            Icons.star,
                            color: Colors.yellow,
                          ),
                          Text(
                            '(5)',
                            style: TextStyle(color: Colors.grey),
                          ),
                        ],
                      ),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(widget.product.name,
                          style: const TextStyle(
                            fontSize: 24,
                            fontWeight: FontWeight.w600,
                            color: Color.fromARGB(255, 89, 87, 87),
                          )),
                      Text("\$${widget.product.price}"),
                    ],
                  ),
                ),
                const Padding(
                  padding: EdgeInsets.all(8.0),
                  child: Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        "Size:",
                        style: TextStyle(
                            color: Colors.grey,
                            fontSize: 20,
                            fontWeight: FontWeight.w500),
                      )),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0),
                  child: SingleChildScrollView(
                    scrollDirection: Axis.horizontal,
                    child: Row(
                      children: [
                        sizeNumber('39'),
                        sizeNumber('40'),
                        sizeNumber('41'),
                        sizeNumber('42'),
                        sizeNumber('43'),
                        sizeNumber('44')
                      ],
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 15.0),
                  child: Text(
                    widget.product.description,
                    style: const TextStyle(color: Color(0xFF666666)),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.only(
                      top: 15.0, bottom: 3, left: 10, right: 10),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    children: [
                      // BlocBuilder<ProductBloc, ProductState>(
                      // builder: (context, state) {
                      // return
                      BlocBuilder<ProductBloc, ProductState>(
                        builder: (context, state) {
                          return OutlinedButton(
                              style: ButtonStyle(
                                  side: WidgetStateProperty.all<BorderSide>(
                                      const BorderSide(
                                    color: Colors.red,
                                    width: 1,
                                  )),
                                  shape: WidgetStateProperty.all<
                                          RoundedRectangleBorder>(
                                      RoundedRectangleBorder(
                                          borderRadius:
                                              BorderRadius.circular(12)))),
                              onPressed: () {
                                context.read<ProductBloc>().add(
                                    DeleteProductEvent(
                                        productId: widget.product.id));

                                Navigator.of(context).pop();
                              },
                              child: const Text("DELETE",
                                  style: TextStyle(color: Colors.red)));
                        },
                      ),
                      // },
                      // ),
                      ElevatedButton(
                        style: ElevatedButton.styleFrom(
                            backgroundColor: Colors.blue[800],
                            shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(8))),
                        onPressed: () {
                          Navigator.of(context).push(MaterialPageRoute(
                              builder: (BuildContext context) {
                            return UpdateProduct(product: widget.product);
                          }));
                        },
                        child: const Text(
                          "UPDATE",
                          style: TextStyle(
                              color: Color.fromARGB(255, 210, 206, 206)),
                        ),
                      )
                    ],
                  ),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
