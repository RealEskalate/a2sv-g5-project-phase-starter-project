import '../widgets/bottom_sheet.dart';
import '../widgets/cards.dart';
import 'pages.dart';

class SearchPage extends StatelessWidget {
  const SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController _searchController = TextEditingController();
    List<ProductEntity> filteredProducts = [];
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back_ios_new),
          onPressed: () {
            context.read<ProductBloc>().add(LoadAllProductEvent());
            Navigator.of(context).pop();
          },
        ),
        title: const Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'Search Product',
              style: TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.w500,
                fontFamily: 'Poppins',
                color: Color(0xff3E3E3E),
              ),
            )
          ],
        ),
      ),
      body: BlocConsumer<ProductBloc, ProductState>(
        listener: (context, state) {},
        builder: (context, state) {
          if (state is ErrorState) {
            return Center(
              child: Text(state.message),
            );
          }
          if (state is LoadingState) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }

          return Padding(
            padding: const EdgeInsets.all(32.0),
            child: Column(
              children: [
                Row(
                  mainAxisSize: MainAxisSize.max,
                  children: [
                    Expanded(
                      child: TextField(
                        controller: _searchController,
                        onChanged: (value) {
                          context
                              .read<ProductBloc>()
                              .add(FilterProductEvent(text: value));
                        },
                        decoration: InputDecoration(
                          labelText: 'Leather',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(8.0),
                          ),
                          suffixIcon: const Icon(Icons.arrow_forward),
                        ),
                      ),
                    ),
                    const SizedBox(width: 8),
                    IconButton(
                      iconSize: 35,
                      icon: const Icon(Icons.filter_alt_sharp,
                          color: Colors.blue),
                      onPressed: () {
                        bottomSheet(context);
                      },
                    ),
                  ],
                ),
                const SizedBox(height: 16),
                Expanded(
                  child: ListView.builder(
                    itemCount:
                        (state as LoadedAllProductsState).products.length,
                    itemBuilder: (context, index) {
                      String _imageUrl = (state).products[index].imageUrl;
                      String _name = (state).products[index].name;
                      dynamic _price = (state).products[index].price;
                      return Cards(
                        imageUrl: _imageUrl,
                        name: _name,
                        price: _price,
                      );
                    },
                  ),
                ),
                const SizedBox(height: 16),
              ],
            ),
          );
        },
      ),
    );
  }
}
