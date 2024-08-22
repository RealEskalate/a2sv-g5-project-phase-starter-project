import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../injection_container.dart';
import '../../domain/entitity/product.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../widgets/size_container.dart';

class ProductDetailPage extends StatelessWidget {
  const ProductDetailPage({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    final Map arguments =
        (ModalRoute.of(context)?.settings.arguments as Map?) ?? {};
    final Product product = arguments['product'];
    return Scaffold(
      body: BlocProvider(
        create: (_) =>
            sl<ProductBloc>()..add(DeleteProductEvent(productId: product.id)),
        child: _ProductDetailPageBody(
          product: product,
        ),
      ),
    );
  }
}

class _ProductDetailPageBody extends StatefulWidget {
  final Product product;

  _ProductDetailPageBody({required this.product});

  @override
  State<_ProductDetailPageBody> createState() => _ProductDetailPageBodyState();
}

class _ProductDetailPageBodyState extends State<_ProductDetailPageBody> {
  int? selectedSize;

  @override
  Widget build(BuildContext context) {
    final Product product = widget.product;
    return BlocProvider(
      create: (_) =>
          sl<ProductBloc>()..add(DeleteProductEvent(productId: product.id)),
      child: SingleChildScrollView(
        child: Column(
          children: [
            Stack(children: [
              AspectRatio(
                aspectRatio: 16 / 12,
                child: Image.network(
                  product.imageUrl,
                  fit: BoxFit.fill,
                ),
              ),
              Container(
                margin:
                    const EdgeInsets.symmetric(horizontal: 10, vertical: 40),
                padding: const EdgeInsets.all(1),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(30),
                  color: Colors.white,
                ),
                child: IconButton(
                  onPressed: () {
                    Navigator.pop(context);
                  },
                  icon: const Icon(Icons.arrow_back_ios_new),
                ),
              ),
            ]),
            const SizedBox(
              height: 10,
            ),
            Container(
              padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 30),
              child: Column(
                children: [
                  const Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        // product.category,
                        'men\s shoe',
                        style: TextStyle(fontSize: 12, color: Colors.grey),
                      ),
                      Row(
                        children: [
                          Icon(
                            Icons.star,
                            size: 15,
                            color: Color.fromARGB(255, 246, 186, 45),
                          ),
                          Text(
                            // product.rating,
                            '4.0',
                            style: TextStyle(fontSize: 12, color: Colors.grey),
                          ),
                        ],
                      ),
                    ],
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        product.name,
                        style: const TextStyle(
                            fontSize: 25, fontWeight: FontWeight.w600),
                      ),
                      Text(product.price.toString()),
                    ],
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  const Row(
                    children: [
                      Text(
                        'Size:',
                        style: TextStyle(fontSize: 20),
                        textAlign: TextAlign.left,
                      ),
                    ],
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  SingleChildScrollView(
                    scrollDirection: Axis.horizontal,
                    child: Row(children: [
                      // for (var sz in product.size)
                      for (var sz in [39, 40, 41, 42, 43, 44])
                        GestureDetector(
                          onTap: () {
                            setState(() {
                              selectedSize = sz;
                            });
                          },
                          child: SizeContainer(
                              size: sz, isSelected: selectedSize == sz),
                        ),
                    ]),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Container(
                      padding: const EdgeInsets.only(bottom: 50),
                      child: Text(
                        product.description,
                        style: const TextStyle(fontSize: 15),
                      )),
                  const SizedBox(
                    height: 10,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      ElevatedButton(
                        onPressed: () async {
                          BlocProvider.of<ProductBloc>(context).add(
                              await DeleteProductEvent(productId: product.id));

                          Navigator.pushNamed(context, '/homepage');
                        },
                        style: ButtonStyle(
                          padding: const WidgetStatePropertyAll(
                              EdgeInsets.symmetric(
                                  vertical: 20, horizontal: 50)),
                          side: WidgetStateProperty.all(
                              const BorderSide(color: Colors.red)),
                          shape: WidgetStatePropertyAll(RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10),
                          )),
                        ),
                        child: const Text('DELETE',
                            style: TextStyle(
                              color: Colors.red,
                            )),
                      ),
                      ElevatedButton(
                        onPressed: () {
                          Navigator.pushNamed(context, '/add_product_page',
                              arguments: {
                                'product': product,
                              });
                        },
                        style: ButtonStyle(
                          padding: const WidgetStatePropertyAll(
                              EdgeInsets.symmetric(
                                  vertical: 20, horizontal: 50)),
                          backgroundColor: WidgetStateProperty.all(
                              const Color.fromARGB(255, 32, 77, 202)),
                          shape: WidgetStatePropertyAll(RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10),
                          )),
                        ),
                        child: const Text('UPDATE',
                            style: TextStyle(color: Colors.white)),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
