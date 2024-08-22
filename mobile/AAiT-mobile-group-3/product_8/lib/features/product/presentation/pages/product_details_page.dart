import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/custom_buttom.dart';

class Detailspage extends StatelessWidget {
  final Product product;
  static const routeName = '/detail';
  const Detailspage({super.key, required this.product});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: BlocListener<ProductBloc, ProductState>(
          listener: (context, state) {
            if (state is ProductDeleteState) {
              ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('deleted successfully')));
              Navigator.of(context).pushNamed('/home');
              // Navigator.pop(context);
            } else if (state is ProductError) {
              ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Error deleting product')));
            }
          },
          child: Column(
            children: [
              Stack(
                children: [
                  SizedBox(
                    height: MediaQuery.of(context).size.height * 0.32,
                    width: double.infinity,
                    child: ClipRRect(
                      borderRadius: const BorderRadius.only(
                          topRight: Radius.circular(20),
                          topLeft: Radius.circular(20)),
                      child: Image.network(
                        product.imageUrl,
                        fit: BoxFit.cover,
                      ),
                    ),
                  ),
                  Positioned(
                    top: 40, // Adjust based on your design
                    left: 10, // Adjust based on your design
                    child: Container(
                      width: 40.0,
                      height: 40.0,
                      decoration: BoxDecoration(
                          color: Colors.white,
                          borderRadius: BorderRadius.circular(20)),
                      // color: Colors.white,
                      child: Center(
                        child: IconButton(
                          // color: Colors.black,
                          onPressed: () {
                            Navigator.pop(context);
                          },
                          icon: Icon(
                            Icons.arrow_back_ios_new_rounded,
                            color: Colors.indigoAccent.shade400,
                          ),
                        ),
                      ),
                    ),
                  ),
                ],
              ),
              Padding(
                padding: const EdgeInsets.all(15.0),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        const Text(
                          "Men's shoe",
                          style:
                              TextStyle(fontSize: 15.0, color: Colors.black45),
                        ),
                        Row(
                          children: [
                            Icon(
                              Icons.star,
                              color: Colors.yellow[700],
                            ),
                            const Text(
                              '(4.0)',
                              style: TextStyle(
                                  fontSize: 15.0, color: Colors.black45),
                            ),
                          ],
                        )
                      ],
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(vertical: 10.0),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(
                            product.name,
                            style: const TextStyle(
                                fontSize: 20.0, fontWeight: FontWeight.bold),
                          ),
                          Text(
                            '\$${product.price}',
                            style: const TextStyle(
                                fontSize: 15.0, fontWeight: FontWeight.bold),
                          )
                        ],
                      ),
                    ),
                    const Text(
                      'Size:',
                      style: TextStyle(fontSize: 20.0),
                    ),
                    SizedBox(
                      height: 60,
                      child: ListView(
                        physics: const AlwaysScrollableScrollPhysics(),
                        scrollDirection: Axis.horizontal,
                        children: [
                          _buildSizeCard('39', '41'),
                          _buildSizeCard('40', '41'),
                          _buildSizeCard('41', '41'),
                          _buildSizeCard('42', '41'),
                          _buildSizeCard('43', '41'),
                          _buildSizeCard('44', '44'),
                          _buildSizeCard('45', '41'),
                        ],
                      ),
                    ),
                    const SizedBox(
                      height: 20.0,
                    ),
                    Text(
                      product.description,
                    ),
                    const SizedBox(
                      height: 20.0,
                    ),
                    BlocBuilder<ProductBloc, ProductState>(
                      builder: (context, state) {
                        return Row(
                          mainAxisAlignment: MainAxisAlignment.spaceBetween,
                          children: [
                            CustomButton(
                                backgroundColor: Colors.white,
                                foregroundColor:
                                    const Color.fromRGBO(255, 19, 19, 0.79),
                                borderColor:
                                    const Color.fromRGBO(255, 19, 19, 0.79),
                                buttonWidth: 120,
                                buttonHeight: 45,
                                child: state is ProductLoading
                                    ? const CircularProgressIndicator(
                                        valueColor: AlwaysStoppedAnimation<
                                                Color>(
                                            Color.fromRGBO(255, 19, 19, 0.79)),
                                      )
                                    : const Text(
                                        'DELETE',
                                        style: TextStyle(
                                            fontWeight: FontWeight.w600),
                                      ),
                                onPressed: () => {
                                      context.read<ProductBloc>().add(
                                          DeleteProductEvent(id: product.id)),
                                      // Navigator.of(context).pushNamed('/home')
                                    }),
                            CustomButton(
                                backgroundColor: Theme.of(context).primaryColor,
                                foregroundColor: Colors.white,
                                borderColor: Theme.of(context).primaryColor,
                                buttonWidth: 120,
                                buttonHeight: 45,
                                child: const Text(
                                  'UPDATE',
                                  style: TextStyle(fontWeight: FontWeight.w600),
                                ),
                                onPressed: () {
                                  Navigator.of(context).pushNamed(
                                      '/detail/update',
                                      arguments: Product(
                                          id: product.id,
                                          imageUrl: product.imageUrl,
                                          name: product.name,
                                          price: product.price,
                                          description: product.description));
                                })
                          ],
                        );
                      },
                    ),
                    const SizedBox(
                      height: 20.0,
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

Widget _buildSizeCard(String size, String _size) {
  bool isSelected = false;
  if (size == _size) {
    isSelected = true;
  }
  return Container(
    color: isSelected ? Colors.indigoAccent.shade400 : Colors.white,
    width: 70,
    height: 60,
    child: Card(
      color: isSelected ? Colors.indigoAccent.shade400 : Colors.white,
      child: Center(
        child: Text(
          size,
          style: TextStyle(
              fontSize: 19,
              fontWeight: FontWeight.bold,
              color: isSelected ? Colors.white : Colors.black),
        ),
      ),
    ),
  );
}
