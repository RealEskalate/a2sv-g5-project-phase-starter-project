import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/utils/dummy_data/products_data.dart';
import '../../../../injection_container.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/loading.dart';
import '../widgets/message_display.dart';
import '../widgets/product_cards.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocProvider(
        create: (_) => sl<ProductBloc>()
          ..add(
              GetAllProductEvent()), // Initialize the bloc and trigger fetching products
        child: SafeArea(
          child: Container(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: Column(
              children: [
                _buildHeader(context),
                const SizedBox(height: 30),
                _buildTitleAndSearch(context),
                const SizedBox(height: 20),
                BlocBuilder<ProductBloc, ProductState>(
                  builder: (context, state) {
                    if (state is ProductStateLoading) {
                      return const LoadingWidget();
                    } else if (state is AllProductsLoaded) {
                      return Expanded(
                        child: ListView.builder(
                          itemCount: state.products.length,
                          itemBuilder: (context, index) {
                            return ProductCard(product: state.products[index]);
                          },
                        ),
                      );
                    } else if (state is AllProductsLoadedFailure) {
                      return MessageDisplay(message: state.message);
                    } else {
                      return const MessageDisplay(message: 'Unknown state');
                    }
                  },
                ),
              ],
            ),
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.pushNamed(context, '/add_product_page',
              arguments: {'products': Products});
        },
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(30)),
        backgroundColor: const Color.fromARGB(255, 33, 75, 243),
        child: const Icon(Icons.add, color: Colors.white),
      ),
    );
  }

  // ... rest of your code
}

Widget _buildHeader(BuildContext context) {
  return Row(
    mainAxisAlignment: MainAxisAlignment.spaceBetween,
    children: [
      Row(
        children: [
          Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
              border: Border.all(color: Colors.grey),
              color: Colors.grey,
            ),
            child: IconButton(
              onPressed: () {},
              icon: const Icon(Icons.person, color: Colors.grey),
            ),
          ),
          const SizedBox(width: 10),
          const Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'August 14, 2024',
                style: TextStyle(fontSize: 10),
              ),
              SizedBox(height: 5),
              Row(
                children: [
                  Text('Hello,'),
                  Text(
                    'Heran',
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                ],
              ),
            ],
          ),
        ],
      ),
      Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          border: Border.all(color: Colors.grey),
        ),
        child: IconButton(
          onPressed: () {},
          icon: const Icon(Icons.notifications_outlined),
        ),
      ),
    ],
  );
}

Widget _buildTitleAndSearch(BuildContext context) {
  return Row(
    mainAxisAlignment: MainAxisAlignment.spaceBetween,
    children: [
      const Text(
        'Available Products',
        style: TextStyle(fontWeight: FontWeight.bold, fontSize: 25),
      ),
      Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          border: Border.all(color: Colors.grey),
        ),
        child: IconButton(
          onPressed: () {
            Navigator.pushNamed(context, '/search_page');
          },
          icon: const Icon(Icons.search, color: Colors.grey),
        ),
      ),
    ],
  );
}

Widget _buildProductList() {
  return BlocBuilder<ProductBloc, ProductState>(
    builder: (context, state) {
      if (state is ProductStateLoading) {
        return const LoadingWidget();
      } else if (state is AllProductsLoaded) {
        return Expanded(
          child: ListView.builder(
            itemCount: state.products.length,
            itemBuilder: (context, index) {
              return ProductCard(product: state.products[index]);
            },
          ),
        );
      } else if (state is AllProductsLoadedFailure) {
        return MessageDisplay(message: state.message);
      } else {
        return const MessageDisplay(message: 'Unknown state');
      }
    },
  );
}


