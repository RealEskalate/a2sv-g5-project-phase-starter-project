import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../injection_container.dart';
import '../../domain/entities/product.dart';
import '../bloc/product_bloc.dart';
import '../widgets/modal_sheet.dart';
import '../widgets/product_card.dart';

class SearchPage extends StatelessWidget {
  const SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Search Product'),
        centerTitle: true,
      ),
      body: Padding(
          padding: const EdgeInsets.all(20.0), child: buildSearch(context)),
    );
  }
}

Widget buildSearch(BuildContext context) {
  context.read<ProductBloc>().add(LoadAllProductEvent());
  return Column(
    children: [
      Row(
        // mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          const SizedBox(
            width: 20,
          ),
          SizedBox(
            width: 280,
            height: 44,
            child: TextField(
              decoration: InputDecoration(
                  contentPadding: const EdgeInsets.only(top: 5, left: 15),
                  suffixIcon: Icon(
                    Icons.arrow_forward,
                    color: Theme.of(context).primaryColor,
                  ),
                  hintText: 'Leather',
                  border: const OutlineInputBorder(
                      borderSide: BorderSide(color: Colors.black12))),
            ),
          ),
          const SizedBox(
            width: 8,
          ),
          GestureDetector(
            onTap: () => showModal(context),
            child: Container(
                decoration: BoxDecoration(
                  color: Theme.of(context).primaryColor,
                  border: Border.all(
                      width: 2, color: Theme.of(context).primaryColor),
                  borderRadius: BorderRadius.circular(5),
                ),
                child: const Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Icon(
                    Icons.filter_list,
                    color: Colors.white,
                  ),
                )),
          ),
          const SizedBox(
            width: 15,
          )
        ],
      ),
      const SizedBox(
        height: 30,
      ),
      BlocBuilder<ProductBloc, ProductState>(builder: (context, state) {
        if (state is ProductLoading) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else if (state is LoadAllProductState) {
          return Expanded(
            child: ListView.builder(
                itemCount: state.products.length,
                itemBuilder: (BuildContext context, int index) {
                  return ProductCard(
                    product: Product(
                        id: state.products[index].id,
                        imageUrl: state.products[index].imageUrl,
                        name: state.products[index].name,
                        price: state.products[index].price,
                        description: state.products[index].description),
                  );
                }),
          );
        } else if (state is ProductErrorState) {
          return Center(
            child: Text('Error: ${state.message}'),
          );
        } else {
          return const Center(
            child: Text('No products available'),
          );
        }
      }),
    ],
  );
}

void showModal(BuildContext context) {
  showModalBottomSheet<void>(
      shape: const RoundedRectangleBorder(borderRadius: BorderRadius.zero),
      context: context,
      builder: (BuildContext context) {
        return const ModalSheet();
      });
}
