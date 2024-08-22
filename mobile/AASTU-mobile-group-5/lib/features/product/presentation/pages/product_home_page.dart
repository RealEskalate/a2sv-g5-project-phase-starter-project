import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../service_locator.dart';
import '../../domain/use_case/get_all_products.dart';
import '../bloc/home_page/home_page_bloc.dart';
import '../widgets/available_header.dart';
import '../widgets/custom_app_bar_home_page.dart';
import '../widgets/item_card.dart';
import '../widgets/search_button.dart';

class HomePage extends StatelessWidget {


  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => HomePageBloc(
        getAllProducts: GetAllProducts(getIt()),
      )..add(FetchAllProductsEvent()),
      child: Scaffold(
        appBar: const CustomAppBar(),
        body: Column(
          children: [
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: Row(
                children: [
                  const Expanded(child: AvailableProductsHeader()),
                  buildSearchButton(context),
                ],
              ),
            ),
            Expanded(
              child: BlocBuilder<HomePageBloc, HomePageState>(
                builder: (context, state) {
                  if (state is HomePageLoadingState) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (state is HomePageLoadedState) {
                    return RefreshIndicator(
                      onRefresh: () async{
                        context.read<HomePageBloc>().add(FetchAllProductsEvent());
                      },
                      child: ListView.builder(
                        itemCount: state.products.length,
                        itemBuilder: (context, index) {
                          return GestureDetector(
                            onTap: () {
                              Navigator.pushNamed(context, '/details',
                                arguments: state.products[index]);
                            },
                            child: Padding(
                              padding: const EdgeInsets.all(8.0),
                              child: ProductItemCard(product: state.products[index]),
                            ));
                        },
                      ),
                    );
                  } else if (state is HomePageErrorState) {
                    return Center(child: Text(state.message));
                  } else {
                    return const Center(child: Text('Failed to load products'));
                  }
                },
              ),
            ),
          ],
        ),
        floatingActionButton: Padding(
          padding: const EdgeInsets.only(bottom: 10.0), // Adjust the value as needed
          child: FloatingActionButton(
            onPressed: () {
              Navigator.pushNamed(context, '/add');
            },
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(30),
            ),
            backgroundColor: const Color.fromARGB(255, 54, 104, 255),
            child: const Icon(
              Icons.add,
              size: 35,
              color: Colors.white,
            ),
          ),
        ),
      ),
    );
  }
}
