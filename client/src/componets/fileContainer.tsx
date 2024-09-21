import { useQuery } from "@tanstack/react-query";
import FilesList from "./fileslist";
const getfile = async () => {
  return fetch("/api/user/getfile", {
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  }).then((res) => res.body);
};
const FileContainer = () => {
  const { isLoading, data } = useQuery({
    queryKey: ["files"],
    // queryFn: () =>
    //   fetch("https://api.github.com/repos/TanStack/query").then((res) =>
    //     res.json()
    //   ),
    queryFn: getfile,
  });
  console.log(isLoading, data?.getReader);
  return (
    <>
      <FilesList fileName="fhf" fileUrl={"fdfsfdfd"} download={downloadFile} />
    </>
  );
};

export default FileContainer;

async function downloadFile(fileurl: URL, fileName: string) {
  try {
    const response = await fetch(fileurl);
    if (!response.ok) throw new Error("Network response was not ok.");

    const blob = await response.blob();

    const url = window.URL.createObjectURL(blob);

    const a = document.createElement("a");
    a.href = url;
    a.download = fileName; // Specify the filename for the download
    document.body.appendChild(a);
    a.click();

    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Error downloading file:", error);
  }
}
