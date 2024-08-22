import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../../../../service_locator.dart';
import '../../domain/use_case/delete_product.dart';
import '../../domain/use_case/get_product.dart';
import '../bloc/details_page/details_page_bloc.dart';
import '../bloc/search_page/search_page_bloc.dart';
import '../widgets/bottom_sheet.dart';
import '../widgets/item_card.dart';
import 'product_details_page.dart'; // Import your product details page

class SearchPage extends StatelessWidget {
  const SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back_ios_new,
              color: Color.fromARGB(255, 54, 104, 255), size: 20),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        title: const Text(
          'Search Product',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.w500,
          ),
        ),
        centerTitle: true,
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      decoration: InputDecoration(
                        suffixIcon: const Icon(Icons.arrow_forward),
                        hintText: 'Search...',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(6),
                        ),
                      ),
                      onChanged: (value) {
                        context.read<SearchPageBloc>().add(SearchProductsEvent(query: value));
                      },
                    ),
                  ),
                  const SizedBox(width: 15),
                  Container(
                    padding: const EdgeInsets.all(8),
                    width: 56,
                    height: 56,
                    decoration: BoxDecoration(
                      color: const Color.fromARGB(255, 54, 104, 255),
                      borderRadius: BorderRadius.circular(6),
                    ),
                    child: IconButton(
                      onPressed: () {
                        showModalBottomSheet(
                          context: context,
                          builder: (context) => const FilterBottomSheet(),
                          shape: const RoundedRectangleBorder(
                            borderRadius: BorderRadius.vertical(
                              top: Radius.circular(16),
                            ),
                          ),
                        );
                      },
                      icon: const Icon(
                        Icons.filter_list,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 10),
            Expanded(
              child: BlocBuilder<SearchPageBloc, SearchPageState>(
                builder: (context, state) {
                  if (state is SearchPageLoadingState) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (state is SearchPageLoadedState) {
                    return RefreshIndicator(
                      onRefresh: () async {
                        // Re-trigger the search with the current query
                        context.read<SearchPageBloc>().add(
                          FetchAllProductsSearchEvent(),
                        );
                      },
                      child: ListView.builder(
                        itemCount: state.products.length,
                        itemBuilder: (context, index) {
                          final product = state.products[index];
                          return GestureDetector(
                            onTap: () {
                              Navigator.push(
                                context,
                                MaterialPageRoute(
                                  builder: (context) => BlocProvider(
                                    create: (context) => DetailsPageBloc(
                                      getIt<GetProduct>(),
                                      getIt<DeleteProduct>(),
                                    )..add(FetchProductByIdEvent(GetProductParams(product.id))),
                                    child: DetailsPage(id: product.id),
                                  ),
                                ),
                              );
                            },
                            child: ProductItemCard(
                              product: product,
                            ),
                            
                          );
                        },
                      ),
                    );
                  } else if (state is SearchPageErrorState) {
                    return Center(child: Text(state.message));
                  }
                  return const SizedBox(height: 1);
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}