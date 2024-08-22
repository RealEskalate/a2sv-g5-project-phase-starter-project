import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../bloc/details_page/details_page_bloc.dart';
import '../bloc/home_page/home_page_bloc.dart';
import '../widgets/custom_back_button.dart';
import '../widgets/delete_button_details.dart';
import '../widgets/size_selector.dart';
import '../widgets/update_button.dart';

class DetailsPage extends StatelessWidget {
  final String id;

  const DetailsPage({super.key, required this.id});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButtonLocation: FloatingActionButtonLocation.startTop,
      floatingActionButton: const CustomBackButton(),
      body: BlocListener<DetailsPageBloc, DetailsPageState>(
        listener: (context, state) {
          if (state is DetailsPageDeletedState) {
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(content: Text('Product deleted successfully')),
            );
            context.read<HomePageBloc>().add(FetchAllProductsEvent());
            Navigator.pushNamed(context, '/home');
          } else if (state is DetailsPageErrorState) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message)),
            );
          }
        },
        child: BlocBuilder<DetailsPageBloc, DetailsPageState>(
          builder: (context, state) {
            if (state is DetailsPageLoadingState) {
              return const Center(child: CircularProgressIndicator());
            } else if (state is DetailsPageLoadedState) {
              final product = state.product;
              return SingleChildScrollView(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Container(
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(10),
                      ),
                      child: Column(
                        children: [
                          // Image Container
                          Container(
                            height: MediaQuery.of(context).size.height * 0.3,
                            width: double.infinity,
                            decoration: const BoxDecoration(
                              borderRadius: BorderRadius.only(
                                topLeft: Radius.circular(10),
                                topRight: Radius.circular(10),
                              ),
                            ),
                            child: ClipRRect(
                              borderRadius: const BorderRadius.only(
                                topLeft: Radius.circular(10),
                                topRight: Radius.circular(10),
                              ),
                              child: Image.network(
                                product
                                    .imageUrl, // Assuming product.imageUrl contains the URL of the image
                                fit: BoxFit.cover,
                                errorBuilder: (context, error, stackTrace) {
                                  return const Icon(Icons.error,
                                      size:
                                          50); // Display an error icon if image loading fails
                                },
                                loadingBuilder:
                                    (context, child, loadingProgress) {
                                  if (loadingProgress == null) return child;
                                  return const Center(
                                    child: CircularProgressIndicator(),
                                  );
                                },
                              ),
                            ),
                          ),
                          // Product Details
                          Padding(
                            padding: const EdgeInsets.all(8.0),
                            child: Container(
                              decoration: const BoxDecoration(
                                borderRadius: BorderRadius.only(
                                  bottomLeft: Radius.circular(10),
                                  bottomRight: Radius.circular(10),
                                ),
                              ),
                              child: Row(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceBetween,
                                children: [
                                  Expanded(
                                    child: Padding(
                                      padding: const EdgeInsets.all(8.0),
                                      child: Column(
                                        crossAxisAlignment:
                                            CrossAxisAlignment.start,
                                        children: [
                                          const SizedBox(height: 10),
                                          Text(
                                            product.name,
                                            style: TextStyle(
                                              fontSize: 20,
                                              color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                                              fontWeight: FontWeight.bold,
                                            ),
                                          ),
                                          const SizedBox(height: 10),
                                          Text(
                                            product.description,
                                            maxLines: 1,
                                            style: TextStyle(
                                              fontSize: 16,
                                              color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                                            ),
                                          ),
                                        ],
                                      ),
                                    ),
                                  ),
                                  Padding(
                                    padding: const EdgeInsets.all(8.0),
                                    child: Column(
                                      children: [
                                        Row(
                                          children: [
                                            const Icon(Icons.star,
                                                color: Colors.yellow, size: 16),
                                            Text(
                                              '(4.0)',
                                              style: TextStyle(
                                                fontSize: 13,
                                                color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                                              ),
                                            ),
                                          ],
                                        ),
                                        const SizedBox(height: 10),
                                        Row(
                                          children: [
                                            Text(
                                              'Price: ${product.price.toString()}',
                                              style: TextStyle(
                                                fontSize: 16,
                                                color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                                              ),
                                            ),
                                          ],
                                        ),
                                      ],
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),

                    Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Text('Size:',
                          style: TextStyle(
                            color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,
                              fontSize: 18, fontWeight: FontWeight.w500)),
                    ),
                    const SizeSelector(), // Replace with SizeSelector widget
                    const SizedBox(height: 16),
                    Padding(
                      padding: const EdgeInsets.all(16.0),
                      child: Text(
                        product.description,
                        style: TextStyle(
                            fontSize: 14,
                            // color: Color.fromRGBO(102, 102, 102, 1)),
                            color: Theme.of(context).brightness == Brightness.dark ? Colors.white : Colors.black,)
                      ),
                    ),
                    const SizedBox(height: 100),
                     Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          const SizedBox(height: 30),
                          DeleteButtonDetails(id: id), // Use DeleteButton widget
                          const SizedBox(height: 30),
                          const SizedBox(width: 50),
                          UpdateButton(
                            product: product,
                          ), // Use UpdateButton widget
                          const SizedBox(height: 30),
                        ],
                      
                    ),
                  ],
                ),
              );
            }
            return const Center(child: Text('Product not found'));
          },
        ),
      ),
    );
  }
}
