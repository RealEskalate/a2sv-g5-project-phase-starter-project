import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/product_card.dart';

class SearchPage extends StatelessWidget {
  const SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          onPressed: () => Navigator.pop(context),
          icon: Icon(
            Icons.arrow_back_ios_outlined,
            color: Colors.indigoAccent.shade400,
          ),
        ),
        title: const Text(
          'Search Product',
          style: TextStyle(fontWeight: FontWeight.bold),
        ),
        centerTitle: true,
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: homeBuilder(context),
      ),
    );
  }
}

Widget homeBuilder(BuildContext context) {
  context.read<ProductBloc>().add(LoadProduct());
  return Column(
    children: [
      Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Container(
            width: 290,
            height: 50,
            decoration: BoxDecoration(
                border: Border.all(color: Colors.black45),
                borderRadius: BorderRadius.circular(10)),
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  const Text(
                    'Leather',
                    style: TextStyle(fontSize: 20, color: Colors.black45),
                  ),
                  Icon(
                    Icons.arrow_forward,
                    color: Colors.indigoAccent.shade400,
                    weight: 30,
                  ),
                ],
              ),
            ),
          ),
          GestureDetector(
            onTap: () => _showBottomSheet(context),
            child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(5),
                color: Colors.indigoAccent.shade400,
              ),
              child: IconButton(
                onPressed: () => _showBottomSheet(context),
                icon: const Icon(
                  Icons.filter_list_rounded,
                  color: Colors.white,
                ),
              ),
            ),
          ),
        ],
      ),
      const SizedBox(
        height: 10,
      ),
      BlocBuilder<ProductBloc, ProductState>(builder: (context, state) {
        // switch (state.runtimeType) {
        if (state is ProductLoading) {
          return const Center(child: CircularProgressIndicator());
        } else if (state is ProductLoaded) {
          final successState = state;
          return Expanded(
            child: RefreshIndicator(
              onRefresh: () async{ context.read<ProductBloc>().add(LoadProduct()); },
              child: GridView.builder(
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 1,
                  crossAxisSpacing: 10,
                  mainAxisSpacing: 10,
                  childAspectRatio: 1.3,
                ),
                itemCount: successState.products.length,
                itemBuilder: (BuildContext context, int index) {
                  return ProductCard(
                    product: successState.products[index],
                  );
                },
              ),
            ),
          );
        } else {
          return const Center(child: Text('Error loading products'));
        }
      }),
    ],
  );
}

void _showBottomSheet(BuildContext context) {
  showModalBottomSheet(
    context: context,
    builder: (BuildContext context) {
      return Container(
        height: 320,
        decoration: const BoxDecoration(color: Colors.white),
        child: Center(
          child: Padding(
            padding: const EdgeInsets.all(20.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                const SizedBox(
                  height: 20,
                ),
                const Text('Category'),
                const SizedBox(
                  height: 10,
                ),
                SizedBox(
                  height: 50,
                  child: TextField(
                    decoration: InputDecoration(
                      border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(10)),
                    ),
                  ),
                ),
                const SizedBox(
                  height: 10,
                ),
                const Text('Price'),
                RangeSlider(
                    values: const RangeValues(0.2, 0.8),
                    activeColor: Colors.indigoAccent.shade400,
                    onChanged: (RangeValues) => {}),
                const SizedBox(
                  height: 30,
                ),
                OutlinedButton(
                    style: OutlinedButton.styleFrom(
                        backgroundColor: Colors.indigoAccent.shade400,
                        foregroundColor: Colors.white,
                        fixedSize: const Size(double.maxFinite, 45),
                        padding: const EdgeInsets.symmetric(
                            horizontal: 50.0, vertical: 15.0),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        )),
                    onPressed: () => {},
                    child: const Text('APPLY',
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                        ))),
              ],
            ),
          ),
        ),
      );
    },
  );
}
